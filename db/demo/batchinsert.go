package demo

import (
	"fmt"

	model "github.com/fatahnuram/sandbox/db"
	"github.com/jmoiron/sqlx"
)

func batchDemoOneDeptTwoEmployees(db *sqlx.DB) error {
	ops := model.Department{Name: "Operations"}

	// insert dept
	opsId, err := model.InsertDepartmentAndReturnId(db, ops)
	if err != nil {
		return err
	}
	ops.Id = opsId

	// define empl array
	empls := [3]model.Employee{
		{Name: "Ihsan", DepId: ops.Id},
		{Name: "Ahmad", DepId: ops.Id},
		{Name: "Bagas", DepId: ops.Id},
	}
	_, err = db.NamedExec(model.NAMEDQUERY_INSERT_EMPL_DEPTID, empls)
	if err != nil {
		return err
	}

	deps, err := model.ListAllDepartments()
	if err != nil {
		return err
	}
	fmt.Printf("All departments: %v\n", deps)
	employees, err := model.ListAllEmployees()
	if err != nil {
		return err
	}
	fmt.Printf("All employees: %v\n", employees)

	return nil
}

func batchDemoAssignEmployee(db *sqlx.DB) error {
	corp, admin := model.Department{Name: "Corp. Strategy"}, model.Department{Name: "Administration"}
	// flow: initially assigned to corp dept, then changed to admin dept

	// set initial state
	corpId, err := model.InsertDepartmentAndReturnId(db, corp)
	if err != nil {
		return err
	}
	corp.Id = corpId

	// create employees
	empls := [3]model.Employee{
		{Name: "Luna", DepId: corp.Id},
		{Name: "Wisnu", DepId: corp.Id},
		{Name: "Satria", DepId: corp.Id},
	}
	_, err = db.NamedExec(model.NAMEDQUERY_INSERT_EMPL_DEPTID, empls)
	if err != nil {
		return err
	}

	deps, err := model.ListAllDepartments()
	if err != nil {
		return err
	}
	fmt.Printf("All departments: %v\n", deps)
	employees, err := model.ListAllEmployees()
	if err != nil {
		return err
	}
	fmt.Printf("All employees: %v\n", employees)

	// create another dept
	adminId, err := model.InsertDepartmentAndReturnId(db, admin)
	if err != nil {
		return err
	}
	admin.Id = adminId

	// assign to other dept
	db.MustExec("update employee set department_id = ? where department_id = ?", admin.Id, corp.Id)

	deps, err = model.ListAllDepartments()
	if err != nil {
		return err
	}
	fmt.Printf("All departments: %v\n", deps)
	employees, err = model.ListAllEmployees()
	if err != nil {
		return err
	}
	fmt.Printf("All employees: %v\n", employees)

	return nil
}
