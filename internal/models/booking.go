package models

import "time"

type Booking struct {
	Id           int
	BookingTable string
	Start        time.Time
	End          time.Time
}
