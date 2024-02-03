package main

import (
	"github.com/jmoiron/sqlx"
)

func batchDemoOneDeptTwoEmployees(db *sqlx.DB) error {
	ops := Department{Name: "Operations"}

	// insert dept
	opsId, err := insertDepartmentAndReturnId(db, ops)
	if err != nil {
		return err
	}
	ops.Id = opsId

	// define empl array
	empls := [3]Employee{
		{Name: "Ihsan", DepId: ops.Id},
		{Name: "Ahmad", DepId: ops.Id},
		{Name: "Bagas", DepId: ops.Id},
	}
	_, err = db.NamedExec(NAMEDQUERY_INSERT_EMPL_DEPTID, empls)
	if err != nil {
		return err
	}

	printAllDept(db)
	printAllEmpl(db)

	return nil
}

func batchDemoAssignEmployee(db *sqlx.DB) error {
	corp, admin := Department{Name: "Corp. Strategy"}, Department{Name: "Administration"}
	// flow: initially assigned to corp dept, then changed to admin dept

	// set initial state
	corpId, err := insertDepartmentAndReturnId(db, corp)
	if err != nil {
		return err
	}
	corp.Id = corpId

	// create employees
	empls := [3]Employee{
		{Name: "Luna", DepId: corp.Id},
		{Name: "Wisnu", DepId: corp.Id},
		{Name: "Satria", DepId: corp.Id},
	}
	_, err = db.NamedExec(NAMEDQUERY_INSERT_EMPL_DEPTID, empls)
	if err != nil {
		return err
	}

	printAllDept(db)
	printAllEmpl(db)

	// create another dept
	adminId, err := insertDepartmentAndReturnId(db, admin)
	if err != nil {
		return err
	}
	admin.Id = adminId

	// assign to other dept
	db.MustExec("update employee set department_id = ? where department_id = ?", admin.Id, corp.Id)

	printAllDept(db)
	printAllEmpl(db)

	return nil
}
