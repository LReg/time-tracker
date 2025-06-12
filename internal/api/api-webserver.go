package api

import (
	"encoding/json"
	"github.com/LReg/time-tracker/internal/timeslot"
	"github.com/LReg/time-tracker/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

func StartTimeSlotApi() {
	mux := http.NewServeMux()
	mux.HandleFunc("/quit", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:49152")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte("OK"))
		if err != nil {
			log.Error("Error writing response", err)
		}
		go func() {
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}()
	})
	mux.HandleFunc("/tt/{ms}", func(writer http.ResponseWriter, request *http.Request) {
		ms := request.PathValue("ms")
		millis, err := strconv.ParseInt(ms, 10, 64)
		if err != nil {
			log.Error("Error converting ms to int", err)
		}

		date := time.UnixMilli(millis)
		from := util.GetLastMidnightFrom(date)
		to := util.GetLastMidnightFrom(date.Add(24 * time.Hour))

		timeslots := timeslot.GetTimeSlotsFromTo(from, to)
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:49152")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if len(timeslots) == 0 {
			writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("Not Found"))
			return
		}

		jsonTimeslots, err := json.Marshal(timeslots)
		if err != nil {
			log.Error("Error converting timeslot to json", err)
			writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error converting timeslot to json"))
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

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
