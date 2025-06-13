package parse

import (
	"github.com/LReg/time-tracker/internal/history"
	"github.com/LReg/time-tracker/internal/output"
	"github.com/LReg/time-tracker/internal/timeslot"
	"os"
	"strings"
)

const (
	ArgHelp               = "help"
	ArgVersion            = "version"
	ArgShowHistory        = "show"
	ArgShowHistoryConsole = "showConsole"
	ArgStart              = "start"
)

func ParseArgsAndHandle() {
	args := os.Args

	if len(args) < 2 {
		output.ShowUsageAndExit()
	}

	if len(args) == 2 {
		switch args[1] {
		case ArgHelp:
			output.ShowHelpAndExit()
		case ArgVersion:
			output.ShowVersionAndExit()
		case ArgShowHistory:
			history.ShowHistoryAndListen()
		case ArgStart:
			timeslot.StartAndExit()
		case ArgShowHistoryConsole:
			history.ShowHistoryConsoleAndExit()
		}
	}

	combined := strings.Join(args[1:], " ")
	timeslot.AppendToPreviousTimeSlotAndExit(combined)
}
