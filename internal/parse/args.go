package parse

import (
	"github.com/LReg/time-tracker/internal/history"
	"github.com/LReg/time-tracker/internal/output"
	"os"
)

const (
	ArgHelp        = "help"
	ArgVersion     = "version"
	ArgShowHistory = "show"
	ArgStart       = "start"
)

func ParseArgsAndHandle() {
	args := os.Args

	if len(args) < 2 {
		output.ShowUsageAndExit()
	}

	if len(args) == 2 {
		switch args[1] {
		case ArgHelp:
		case ArgVersion:
		case ArgShowHistory:
			history.ShowHistory()
		case ArgStart:
		}
	}

}
