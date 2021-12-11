package handlers

import (
	"github.com/vladimish/client-bot/internal/db"
	"github.com/vladimish/client-bot/internal/server/requests"
)

func HandleAddTable(req requests.AddTable) error {
	return db.GetDB().AddTable(req.Name)
}
