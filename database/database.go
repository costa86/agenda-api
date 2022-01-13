package database

import (
	"agenda/helpers"
	"agenda/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	helpers.CheckError(err)

	db.AutoMigrate(&models.Person{})
	// db.AutoMigrate(&models.Stack{})
	db.AutoMigrate(&models.TimeSlot{})

	helpers.Print("Connected to db")
	DB = db
}

func AddPerson(personSquema helpers.PersonSchema) (models.Person, string) {
	person := models.Person{Name: personSquema.Name, Profile: personSquema.Profile}

	var timeSlotList []time.Time

	if personSquema.Profile != helpers.Candidate && personSquema.Profile != helpers.Interviewer {
		return models.Person{}, "invalid profile"
	}

	for _, i := range personSquema.TimeSlots {
		time, timeIsValid := helpers.CastStringToTime(i)

		if !timeIsValid {
			return models.Person{}, "invalid time slot"
		}
		timeSlotList = append(timeSlotList, time)
	}

	DB.Create(&person)

	for _, i := range timeSlotList {
		DB.Create(&models.TimeSlot{Slot: i, PersonID: person.ID})
	}
	return person, helpers.Ok

}

func GetPersonByID(id string) (models.Person, bool) {
	var person models.Person

	DB.First(&person, id)

	if person.Name == "" {
		return person, false
	}
	return person, true
}

func DeletePersonTimeSlots(person models.Person) {
	personTimeSlots := GetPersonTimeSlots(person)
	for _, v := range personTimeSlots {
		DB.Delete(&v, v.ID)
	}
}

func GetPersonTimeSlots(person models.Person) []models.TimeSlot {
	var personTimeSlotList []models.TimeSlot
	DB.Find(&personTimeSlotList, "person_id = ?", person.ID)
	return personTimeSlotList
}

func AddTimeSlotsToPerson(timeSlotList []time.Time, person models.Person) {
	for _, v := range timeSlotList {
		DB.Create(&models.TimeSlot{Slot: v, PersonID: person.ID})
	}
}

func GetPersonStringTimeSlots(person models.Person) []string {
	var timeSlotList []models.TimeSlot
	var timeSlotListString []string

	DB.Find(&timeSlotList, "person_id = ?", person.ID)
	for _, i := range timeSlotList {
		timeSlotListString = append(timeSlotListString, i.Slot.String())
	}
	return timeSlotListString
}

func UpdatePersonSchema(personSchema helpers.PersonSchema, person models.Person) bool {
	if personSchema.Profile != helpers.KeywordDoNotChangeField {
		if personSchema.Profile != helpers.Candidate && personSchema.Profile != helpers.Interviewer {
			return false
		}
		DB.Model(&person).Update("Profile", personSchema.Profile)
	}
	return true
}

func UpdatePersonTimeSlots(personSchema helpers.PersonSchema, person models.Person) bool {
	if len(personSchema.TimeSlots) > 0 {
		timeSlotList, timeSlotListIsValid := helpers.CastStringListToTimeList(personSchema.TimeSlots)
		if !timeSlotListIsValid {
			return false
		}
		DeletePersonTimeSlots(person)
		AddTimeSlotsToPerson(timeSlotList, person)
	}
	return true
}
