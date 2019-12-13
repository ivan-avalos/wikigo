package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DBSettings represents DB connection parameters
type DBSettings struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

func (dbs *DBSettings) initDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbs.Username+":"+dbs.Password+"@/"+dbs.Database)
	if err != nil {
		return nil, err
	}
	return db, nil
}
