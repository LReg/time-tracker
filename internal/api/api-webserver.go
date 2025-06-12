package api

import (
	"encoding/json"
	"github.com/LReg/time-tracker/internal/timeslot"
	"github.com/LReg/time-tracker/internal/ui"
	"github.com/LReg/time-tracker/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

func setCorsHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", ui.WebserverUrl)
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func notFound(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("Not Found"))
}

func internalServerError(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte("Error converting timeslot to json"))
}

func writeJson(writer http.ResponseWriter, err error, jsonTimeslots []byte) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = writer.Write(jsonTimeslots)
	return err
}

func StartTimeSlotApi() {
	mux := http.NewServeMux()

	mux.HandleFunc("/quit", func(writer http.ResponseWriter, request *http.Request) {
		setCorsHeaders(writer)
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
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
		from, to := parseMS(request)
		timeslots := timeslot.GetTimeSlotsFromTo(from, to)

		setCorsHeaders(writer)

		if len(timeslots) == 0 {
			notFound(writer)
			return
		}

		jsonTimeslots, err := json.Marshal(timeslots)
		if err != nil {
			log.Error("Error converting timeslot to json", err)
			internalServerError(writer)
			return
		}

		err = writeJson(writer, err, jsonTimeslots)
		if err != nil {
			log.Error("Error writing response", err)
		}
	})
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal("Could not start the webserver", err)
	}
}

func parseMS(request *http.Request) (time.Time, time.Time) {
	ms := request.PathValue("ms")
	millis, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		log.Error("Error converting ms to int", err)
	}

	date := time.UnixMilli(millis)
	from := util.GetLastMidnightFrom(date)
	to := util.GetLastMidnightFrom(date.Add(24 * time.Hour))
	return from, to
}
