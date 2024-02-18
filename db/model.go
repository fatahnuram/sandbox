package db

import (
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

func InsertDepartmentAndReturnId(db *sqlx.DB, dep Department) (int64, error) {
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

func InsertEmployeeAndReturnId(db *sqlx.DB, empl Employee) (int64, error) {
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

func GetDepartmentById(id int64) (Department, error) {
	db, err := GetDBConnection()
	if err != nil {
		return Department{}, err
	}

	var dep Department
	err = db.Get(&dep, "select * from department where id = ?", id)
	if err != nil {
		return Department{}, err
	}

	return dep, nil
}

func GetEmployeeById(id int64) (Employee, error) {
	db, err := GetDBConnection()
	if err != nil {
		return Employee{}, err
	}

	var empl Employee
	err = db.Get(&empl, "select * from employee where id = ?", id)
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

func deleteAndReturnDepartmentById(id int64) (Department, error) {
	dep, err := GetDepartmentById(id)
	if err != nil {
		return Department{}, err
	}
	db, err := GetDBConnection()
	if err != nil {
		return Department{}, err
	}
	db.MustExec("delete from department where id = ?", id)
	return dep, nil
}

func DeleteEmployeeById(id int64) (int64, error) {
	db, err := GetDBConnection()
	if err != nil {
		return -1, err
	}
	result := db.MustExec("delete from employee where id = ?", id)
	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rows, nil
}

func deleteAndReturnEmployeeById(id int64) (Employee, error) {
	empl, err := GetEmployeeById(id)
	if err != nil {
		return Employee{}, err
	}
	db, err := GetDBConnection()
	if err != nil {
		return Employee{}, err
	}
	db.MustExec("delete from employee where id = ?", id)
	return empl, nil
}

func ListAllDepartments() ([]Department, error) {
	var deps []Department
	db, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	if err = db.Select(&deps, "select * from department"); err != nil {
		return nil, err
	}
	return deps, nil
}

func ListAllEmployees() ([]Employee, error) {
	var empls []Employee
	db, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	if err := db.Select(&empls, "select * from employee"); err != nil {
		return nil, err
	}
	return empls, nil
}
