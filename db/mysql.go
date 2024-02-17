package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

	return db, nil
}
