package stormdb

import (
	"github.com/asdine/storm/v3"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var DB *storm.DB

func ConnectToDB() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Fehler beim Ermitteln des Binary-Pfads:", err)
		return
	}
	exeDir := filepath.Dir(exePath)
	db, err := storm.Open(exeDir + "/timetracker.db")
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	DB = db
}
