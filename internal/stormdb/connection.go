package stormdb

import (
	"github.com/asdine/storm/v3"
	log "github.com/sirupsen/logrus"
)

var DB *storm.DB

func ConnectToDB() {
	db, err := storm.Open("./timetracker.db")
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	DB = db
}
