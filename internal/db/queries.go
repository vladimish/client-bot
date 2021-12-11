package db

import (
	"fmt"
	"github.com/vladimish/client-bot/internal/models"
	"github.com/vladimish/client-bot/pkg/log"
)

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
	err = rows.Close()
	if err != nil {
		log.Get().Warning(err)
	}

	return tables, nil
}

func (db *DB) ChangeUserState(userId int64, state string, stateData string) error {
	query := fmt.Sprintf("INSERT INTO states (user_id, state, state_data) VALUES (%d, '%s', '%s') ON DUPLICATE KEY UPDATE user_id=%d, state='%s', state_data='%s';", userId, state, stateData, userId, state, stateData)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) GetUserState(userId int64) (state string, stateData string, err error) {
	query := fmt.Sprintf("SELECT `state`, `state_data` FROM states WHERE user_id = %d;", userId)
	err = db.db.QueryRow(query).Scan(&state, &stateData)
	return state, stateData, err
}
