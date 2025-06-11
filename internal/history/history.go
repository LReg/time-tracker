package history

import (
	"encoding/json"
	"fmt"
	"github.com/LReg/time-tracker/internal/models"
	"github.com/LReg/time-tracker/internal/timeslot"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func ShowHistory() {
	fmt.Println("Not implemented yet")
	os.Exit(1)
}

func ShowHistoryConsole() {
	for _, ts := range getHistoryForTime(time.Now()) {
		marshal, err := json.Marshal(ts)
		if err != nil {
			log.Error("Error marshalling history: ", err)
		}
		fmt.Println(string(marshal))
	}
}

func getHistoryForTime(time time.Time) []models.TimeSlot {
	return timeslot.GetTimeSlotsFromTo(getLastMidnightFrom(time), time)
}

func getLastMidnightFrom(from time.Time) time.Time {
	year, month, day := from.Date()
	lastMidnight := time.Date(year, month, day, 0, 0, 0, 0, from.Location())
	return lastMidnight
}
