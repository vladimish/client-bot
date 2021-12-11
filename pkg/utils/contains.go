package utils

import "github.com/vladimish/client-bot/internal/models"

func ContainsTable(tables []models.Table, tableName string) bool {
	for i := range tables {
		if tables[i].Name == tableName {
			return false
		}
	}
	return true
}
