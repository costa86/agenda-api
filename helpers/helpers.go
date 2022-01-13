package helpers

import (
	"fmt"
	"time"
)

var Print = fmt.Println

const (
	TimeSlotFormat          = "2006-01-02T15"
	Ok                      = "ok"
	KeywordDeleteTimeSlots  = "clear"
	KeywordDoNotChangeField = "0"
	Candidate               = "candidate"
	Interviewer             = "interviewer"
	Port                    = "8080"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CastStringToTime(rawTimeSlot string) (time.Time, bool) {

	cleanedTimeSlot, err := time.Parse(TimeSlotFormat, rawTimeSlot)

	if err == nil {
		return cleanedTimeSlot, true
	}

	return time.Time{}, false

}

type PersonSchema struct {
	TimeSlots []string `binding:"required"`
	Profile   string   `binding:"required"`
	Name      string   `binding:"required"`
}

func CastStringListToTimeList(timeSlotsString []string) ([]time.Time, bool) {
	var timeSlotList []time.Time

	for _, v := range timeSlotsString {
		if v == KeywordDeleteTimeSlots {
			return timeSlotList, true
		}
		time, timeIsValid := CastStringToTime(v)

		if !timeIsValid {
			return timeSlotList, false
		}
		timeSlotList = append(timeSlotList, time)
	}
	return timeSlotList, true

}
