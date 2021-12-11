package db

import (
	"fmt"
	"github.com/vladimish/client-bot/internal/models"
	"github.com/vladimish/client-bot/pkg/log"
	"time"
)

func (db *DB) GetAllTables() ([]models.Table, error) {
	query := "SELECT `id`, `table_name` FROM tables;"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	tables := make([]models.Table, 0)

	for rows.Next() {
		table := models.Table{}
		err := rows.Scan(&table.Id, &table.Name)
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

func (db *DB) CreateBookingCallback(userId int64, tableId int) error {
	query1 := fmt.Sprintf("DELETE FROM confirmation_callbacks WHERE user_id=%d;", userId)
	query2 := fmt.Sprintf("INSERT INTO confirmation_callbacks (user_id, table_id) VALUES (%d, %d);", userId, tableId)
	_, err := db.db.Exec(query1)
	if err != nil {
		return err
	}
	_, err = db.db.Exec(query2)
	return err
}

func (db *DB) GetBookingCallback(userId int64) (tableId int, err error) {
	query := fmt.Sprintf("SELECT table_id FROM confirmation_callbacks WHERE user_id=%d;", userId)
	err = db.db.QueryRow(query).Scan(&tableId)
	return tableId, err
}

func (db *DB) GetAllBookings(tableName string) ([]models.Booking, error) {
	res := make([]models.Booking, 0)
	query := fmt.Sprintf("SELECT id, booking_table, start, end FROM booking WHERE booking_table='%s';", tableName)
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var btid string
		var id int
		var start, end int64
		err = rows.Scan(&id, &btid, &start, &end)
		if err != nil {
			return nil, err
		}
		res = append(res, models.Booking{
			Id:           id,
			BookingTable: btid,
			Start:        time.Unix(start, 0),
			End:          time.Unix(end, 0),
		})
	}
	err = rows.Close()
	if err != nil {
		log.Get().Warning(err)
	}

	return res, nil
}

func (db *DB) SaveLastBooked(userId int64, tableName string) error {
	query := fmt.Sprintf("INSERT INTO last_booked (last_user_id, table_name) VALUES (%d, '%s') ON DUPLICATE KEY UPDATE last_user_id=%d, table_name='%s';", userId, tableName, userId, tableName)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) GetLastBooked(userId int64) (tableName string, err error) {
	query := fmt.Sprintf("SELECT table_name FROM last_booked WHERE last_user_id=%d;", userId)
	err = db.db.QueryRow(query).Scan(&tableName)
	return tableName, err
}

func (db *DB) SaveBooking(userId int64, tableName string, start time.Time, end time.Time) error {
	query := fmt.Sprintf("INSERT INTO booking (user_id, booking_table, start, end) VALUES(%d, '%s', %d, %d);", userId, tableName, start.Unix(), end.Unix())
	_, err := db.db.Exec(query)
	return err
}
