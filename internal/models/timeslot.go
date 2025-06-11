package models

import "time"

type TimeSlot struct {
	ID       int // primary key
	Text     string
	MetaJson string
	Start    time.Time `storm:"index"`
	End      time.Time `storm:"index"`
}
