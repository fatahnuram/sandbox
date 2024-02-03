package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	NAMEDQUERY_INSERT_DEPT        = "insert into department (name) values (:name)"
	NAMEDQUERY_INSERT_EMPL_DEPTID = "insert into employee (name, department_id) values (:name, :department_id)"
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
	res, err := db.NamedExec(NAMEDQUERY_INSERT_DEPT, dep)
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
	res, err := db.NamedExec(NAMEDQUERY_INSERT_EMPL_DEPTID, empl)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getDepartmentById(db *sqlx.DB, id int64) (Department, error) {
	var dep Department
	err := db.Get(&dep, "select * from department where id = ?", id)
	if err != nil {
		return Department{}, err
	}
	return dep, nil
}

func getEmployeeById(db *sqlx.DB, id int64) (Employee, error) {
	var empl Employee
	err := db.Get(&empl, "select * from employee where id = ?", id)
	if err != nil {
		return Employee{}, err
	}
	return empl, nil
}

func updateDepartmentNameById(db *sqlx.DB, id int64, name string) {
	db.MustExec("update department set name = ? where id = ?", name, id)
}

func updateEmployeeNameById(db *sqlx.DB, id int64, name string) {
	db.MustExec("update employee set name = ? where id = ?", name, id)
}

func deleteAndReturnDepartmentById(db *sqlx.DB, id int64) (Department, error) {
	dep, err := getDepartmentById(db, id)
	if err != nil {
		return Department{}, err
	}
	db.MustExec("delete from department where id = ?", id)
	return dep, nil
}

func deleteAndReturnEmployeeById(db *sqlx.DB, id int64) (Employee, error) {
	empl, err := getEmployeeById(db, id)
	if err != nil {
		return Employee{}, err
	}
	db.MustExec("delete from employee where id = ?", id)
	return empl, nil
}

func printAllDept(db *sqlx.DB) error {
	var deps []Department
	if err := db.Select(&deps, "select * from department"); err != nil {
		return err
	}
	fmt.Printf("All departments: %v\n", deps)
	return nil
}

func printAllEmpl(db *sqlx.DB) error {
	var empls []Employee
	if err := db.Select(&empls, "select * from employee"); err != nil {
		return err
	}
	fmt.Printf("All employees: %v\n", empls)
	return nil
}
