package tt

import (
	"github.com/vladimish/client-bot/internal/db"
	"time"
)

func IsValid(tableName string, start time.Time, end time.Time) (bool, error) {
	bookings, err := db.GetDB().GetAllBookings(tableName)
	if err != nil {
		return false, err
	}

	for i := range bookings {
		current := bookings[i]
		if start.After(current.Start) && start.Before(current.End) || end.Before(current.End) && end.After(current.Start) {
			return false, nil
		}
	}

	return true, nil
}
