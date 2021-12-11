package handlers

import (
	"github.com/vladimish/client-bot/internal/db"
	"github.com/vladimish/client-bot/internal/server/requests"
)

func HandleDeleteTable(req requests.DeleteTable) error {
	return db.GetDB().DeleteTable(req.Name)
}
