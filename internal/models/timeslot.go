package models

import (
	"fmt"
	"time"
)

type TimeSlot struct {
	ID       int `storm:"id,increment"`
	Text     string
	MetaJson string
	Start    time.Time `storm:"index"`
	End      time.Time `storm:"index"`
}

func (ts TimeSlot) Print() {
	const layout = "15:04:05" // Go's reference time layout

	fmt.Println("──── TimeSlot ─────────────────────────")
	fmt.Printf("Text     : %s\n", ts.Text)
	fmt.Printf("Start    : %s\n", ts.Start.Format(layout))
	fmt.Printf("End      : %s\n", ts.End.Format(layout))
	fmt.Println("──────────────────────────────────────")
}
