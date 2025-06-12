package ui

import (
	"embed"
	"github.com/LReg/time-tracker/internal/util"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"time"
)

var WebserverUrl string = "http://localhost:49152"

//go:embed web/*
var uiFiles embed.FS

func StartStaticWebServer() {
	subFS, err := fs.Sub(uiFiles, "web")
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(subFS)))
	go func() {
		time.Sleep(1 * time.Second)
		util.OpenBrowser(WebserverUrl)
	}()
	err = http.ListenAndServe(":49152", mux)
	if err != nil {
		log.Fatal("Could not start webserver", err)
	}
}
