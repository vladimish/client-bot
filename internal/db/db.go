package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vladimish/client-bot/internal/cfg"
)

type DB struct {
	db *sql.DB
}

var instance *DB

func GetDB() *DB {
	if instance == nil {
		var err error
		instance, err = newDB()
		if err != nil {
			panic(err)
		}
	}
	return instance
}

func newDB() (*DB, error) {
	db, err := sql.Open("mysql", cfg.GetConfig().DBUser+":"+cfg.GetConfig().DBPass+"@tcp("+cfg.GetConfig().DBAddr+")"+"/db")
	if err != nil {
		return nil, err
	}

	bd := &DB{
		db: db,
	}

	return bd, nil
}
