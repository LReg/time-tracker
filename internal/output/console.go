package output

import (
	"fmt"
	"github.com/LReg/time-tracker/internal/version"
	"os"
)

func ShowUsageAndExit() {
	fmt.Println("Usage: tt [start, show, <anything you did>]. Use tt help for more information.")
	os.Exit(1)
}

func ShowVersionAndExit() {
	fmt.Printf("Version: %s\nCommit: %s\nDate: %s\n", version.Version, version.Commit, version.Date)
	os.Exit(0)
}

func ShowHelpAndExit() {
	fmt.Printf("Time Tracker Help\ntt [arg]\nArgs List:\n	start - start the tracking for your day. The information about what you did will be set with the next execution\n	show - will display a gui in the browser to visualize your day(s).\n	[any string] - Will start a new timeslot and put the provided string as what you did since the last execution")
	os.Exit(0)
}
