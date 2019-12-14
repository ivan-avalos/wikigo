package main

import (
	"github.com/jinzhu/gorm"

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

func (dbs *DBSettings) initDb() (*gorm.DB, error) {
	// user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", dbs.Username+":"+dbs.Password+"@/"+dbs.Database+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
