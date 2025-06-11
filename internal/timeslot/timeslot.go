package timeslot

import (
	"github.com/LReg/time-tracker/internal/models"
	"github.com/LReg/time-tracker/internal/stormdb"
	"github.com/asdine/storm/v3/q"
	log "github.com/sirupsen/logrus"
	"time"
)

func Start() {
	timeslot := models.TimeSlot{
		Text:  "Start new episode",
		Start: time.Now(),
		End:   time.Now(),
	}
	err := stormdb.DB.Save(&timeslot)
	if err != nil {
		log.Fatal("Start timeslot could not be saved to DB: ", err)
	}
}

func AppendToPreviousTimeSlot(text string) {
	var slot models.TimeSlot
	err := stormdb.DB.Select().
		OrderBy("End").
		Reverse().
		Limit(1).
		First(&slot)
	if err != nil {
		log.Fatal("Could not find start slot", err)
	}
	nextTimeSlot := models.TimeSlot{
		Text:  text,
		Start: slot.End,
		End:   time.Now(),
	}
	err = stormdb.DB.Save(&nextTimeSlot)
	if err != nil {
		log.Fatal("Append to previous time slot could not be saved to DB: ", err)
	}
}

func GetTimeSlotsFromTo(from time.Time, to time.Time) []models.TimeSlot {
	var slots []models.TimeSlot
	err := stormdb.DB.Select(q.And(q.Gte("Start", from), q.Lte("End", to))).OrderBy("End").Find(&slots)
	if err != nil {
		log.Fatal("Could not find time slots from to: ", err)
	}
	return slots
}
