package main

import (
	"github.com/jmoiron/sqlx"
)

type Department struct {
	Id   int64
	Name string
}

type Employee struct {
	Id    int64
	Name  string
	DepId int64 `db:"department_id"`
}

func insertDepartmentAndReturnId(db *sqlx.DB, dep Department) (int64, error) {
	res, err := db.NamedExec("insert into department (name) values (:name)", dep)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func insertEmployeeAndReturnId(db *sqlx.DB, empl Employee) (int64, error) {
	res, err := db.NamedExec("insert into employee (name, department_id) values (:name, :department_id)", empl)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
