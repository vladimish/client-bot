package handlers

import (
	"github.com/vladimish/client-bot/internal/db"
	"github.com/vladimish/client-bot/internal/server/models"
)

func HandleGetTables() ([]models.Table, error) {
	res, err := db.GetDB().GetAllTablesWithBooking()
	if err != nil {
		return nil, err
	}

	final := make([]models.Table, 0)

	for i := range res {
		final = append(final, models.Table{
			Name:   i,
			Starts: res[i][0],
			Ends:   res[i][1],
		})
	}

	return final, nil
}
