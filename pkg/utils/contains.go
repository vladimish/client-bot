package utils

import "github.com/vladimish/client-bot/internal/models"

func ContainsTable(tables []models.Table, tableName string) int {
	for i := range tables {
		if tables[i].Name == tableName {
			return tables[i].Id
		}
	}
	return -1
}
