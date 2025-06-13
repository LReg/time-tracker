package history

import (
	"github.com/LReg/time-tracker/internal/api"
	"github.com/LReg/time-tracker/internal/models"
	"github.com/LReg/time-tracker/internal/timeslot"
	"github.com/LReg/time-tracker/internal/ui"
	"github.com/LReg/time-tracker/internal/util"
	"os"
	"time"
)

func ShowHistoryAndListen() {
	go api.StartTimeSlotApi()
	ui.StartStaticWebServer()
	os.Exit(1)
}

func ShowHistoryConsoleAndExit() {
	for _, ts := range getHistoryForTime(time.Now()) {
		ts.Print()
	}
	os.Exit(0)
}

func getHistoryForTime(time time.Time) []models.TimeSlot {
	return timeslot.GetTimeSlotsFromTo(util.GetLastMidnightFrom(time), time)
}
