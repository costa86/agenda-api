package routes

import (
	"agenda/database"
	"agenda/helpers"
	"agenda/models"
	"fmt"

	"github.com/ghetzel/go-stockutil/sliceutil"
	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var personSquema helpers.PersonSchema

	if c.BindJSON(&personSquema) == nil {

		person, response := database.AddPerson(personSquema)

		if response != helpers.Ok {
			c.JSON(200, gin.H{"error": response})
			return
		}
		successMessage := fmt.Sprintf("person created. ID: %d, Name:%s", person.ID, person.Name)
		c.JSON(200, gin.H{"success": successMessage})
		return

	}
	c.JSON(200, gin.H{"error": "invalid schema"})

}

func GetPeople(c *gin.Context) {
	var people []models.Person
	database.DB.Find(&people)
	c.JSON(200, people)
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")
	person, personIsFound := database.GetPersonByID(id)

	if !personIsFound {
		c.JSON(200, gin.H{"error": "person ID " + id + " was not found"})
		return
	}
	personTimeSlots := database.GetPersonStringTimeSlots(person)

	if len(personTimeSlots) > 0 {
		c.JSON(200, gin.H{"Name": person.Name, "Profile": person.Profile, "TimeSlots": personTimeSlots})
		return
	}
	c.JSON(200, gin.H{"Name": person.Name, "Profile": person.Profile})
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")

	var personSquema helpers.PersonSchema

	if c.BindJSON(&personSquema) != nil {
		c.JSON(200, gin.H{"error": "invalid schema"})
		return
	}

	person, personIsFound := database.GetPersonByID(id)

	if !personIsFound {
		c.JSON(200, gin.H{"error": "person ID " + id + " was not found"})
		return
	}

	if personSquema.Name != helpers.KeywordDoNotChangeField {
		database.DB.Model(&person).Update("Name", personSquema.Name)
	}

	if !database.UpdatePersonSchema(personSquema, person) {
		c.JSON(200, gin.H{"error": "invalid profile"})
		return
	}
	if !database.UpdatePersonTimeSlots(personSquema, person) {
		c.JSON(200, gin.H{"error": "invalid time slot"})
		return
	}

	successMessage := fmt.Sprintf("person ID %d has been updated", person.ID)
	c.JSON(200, gin.H{"success": successMessage})

}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	person, personIsFound := database.GetPersonByID(id)

	if !personIsFound {
		c.JSON(200, gin.H{"error": "person ID " + id + " was not found"})
		return
	}
	database.DB.Delete(&person, id)
	c.JSON(200, gin.H{"error": "person ID " + id + " has been deleted"})
}

func GetSlots(c *gin.Context) {
	peopleIDs := c.Request.URL.Query()["id"]
	var peopleSharedSlots []interface{}
	var firstPerson models.Person

	for idx, id := range peopleIDs {
		person, personIsFound := database.GetPersonByID(id)
		if !personIsFound {
			c.JSON(200, gin.H{"error": "person ID " + id + " was not found"})
			return
		}
		if idx == 0 {
			firstPerson = person
		}
		currentPersonTimeSlots := database.GetPersonStringTimeSlots(person)
		firstPersonTimeSlots := database.GetPersonStringTimeSlots(firstPerson)
		peopleSharedSlots = sliceutil.Intersect(currentPersonTimeSlots, firstPersonTimeSlots)
	}
	c.JSON(200, peopleSharedSlots)

}
