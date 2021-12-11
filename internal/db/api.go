package db

import "fmt"

func (db *DB) AddTable(name string) error {
	query := fmt.Sprintf("INSERT INTO tables (table_name) VALUES('%s');", name)
	_, err := db.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteTable(name string) error {
	query := fmt.Sprintf("DELETE FROM tables WHERE table_name='%s';", name)
	_, err := db.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetAllTablesWithBooking() (map[string][][]int64, error) {
	res := make(map[string][][]int64)
	query := fmt.Sprintf("SELECT table_name, start, end FROM tables LEFT JOIN booking ON table_name = booking_table;")
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var name string
		var start, end int64
		err = rows.Scan(&name, &start, &end)
		if err != nil {
			return nil, err
		}
		if res[name] == nil {
			res[name] = make([][]int64, 2)
			res[name][0] = make([]int64, 1)
			res[name][0][0] = start
			res[name][1] = make([]int64, 1)
			res[name][1][0] = end
		} else {
			res[name][0] = append(res[name][0], start)
			res[name][1] = append(res[name][1], end)
		}
	}
	return res, nil
}
