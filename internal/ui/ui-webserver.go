package ui

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartStaticWebServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../web")))
	err := http.ListenAndServe(":4200", mux)
	if err != nil {
		log.Fatal("Could not start webserver", err)
	}
}
