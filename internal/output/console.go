package output

import (
	"fmt"
	"os"
)

func ShowUsageAndExit() {
	fmt.Println("Usage: tt [start, show, <anything you did>]. Use tt help for more information.")
	os.Exit(1)
}

func ShowVersionAndExit() {

}
