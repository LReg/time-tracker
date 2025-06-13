package main

import (
	"github.com/LReg/time-tracker/internal/parse"
	"github.com/LReg/time-tracker/internal/stormdb"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func main() {
	stormdb.ConnectToDB()
	parse.ParseArgsAndHandle()
}
