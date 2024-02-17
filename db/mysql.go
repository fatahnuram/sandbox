package db

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DbConn *sqlx.DB = nil

func InitializeMysql(host, user, password, dbName, netProto string) (*sqlx.DB, error) {
	config := mysql.Config{
		User:   user,
		Passwd: password,
		DBName: dbName,
		Net:    netProto,
		Addr:   host,
	}

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	DbConn = db
	return db, nil
}

func GetDBConnection() (*sqlx.DB, error) {
	if DbConn == nil {
		return nil, errors.New("database not initialized")
	}

	return DbConn, nil
}
