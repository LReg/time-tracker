package api

import (
	"encoding/json"
	"fmt"
	"github.com/LReg/time-tracker/internal/timeslot"
	"github.com/LReg/time-tracker/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func StartTimeSlotApi() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tt/{ms}", func(writer http.ResponseWriter, request *http.Request) {
		ms := request.PathValue("ms")
		millis, err := strconv.ParseInt(ms, 10, 64)
		if err != nil {
			log.Error("Error converting ms to int", err)
		}
		fmt.Println("millis", millis)

		date := time.UnixMilli(millis)
		from := util.GetLastMidnightFrom(date)
		to := util.GetLastMidnightFrom(date.Add(24 * time.Hour))

		fmt.Println("from", from, "to", to)

		timeslots := timeslot.GetTimeSlotsFromTo(from, to)

		fmt.Println("timeslots", timeslots)
		if len(timeslots) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Not Found"))
			return
		}

		jsonTimeslots, err := json.Marshal(timeslots)
		if err != nil {
			log.Error("Error converting timeslot to json", err)
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error converting timeslot to json"))
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		_, err = writer.Write(jsonTimeslots)
		if err != nil {
			log.Error("Error writing response", err)
		}
	})
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal("Could not start the webserver", err)
	}
}
