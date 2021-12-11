package db

import "github.com/vladimish/client-bot/internal/models"

func (db *DB) GetAllTables() ([]models.Table, error) {
	query := "SELECT `table_name` FROM tables;"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	tables := make([]models.Table, 0)

	for rows.Next() {
		table := models.Table{}
		err := rows.Scan(&table.Name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}
