package main

import (
	"github.com/LReg/time-tracker/internal/parse"
	"github.com/LReg/time-tracker/internal/stormdb"
)

func main() {
	stormdb.ConnectToDB()
	parse.ParseArgsAndHandle()
}
