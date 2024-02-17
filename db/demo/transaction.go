package demo

import (
	"fmt"

	model "github.com/fatahnuram/sandbox/db"
	"github.com/jmoiron/sqlx"
)

func txOneDeptTwoEmpl(db *sqlx.DB) error {
	// define obj to insert

	sales := model.Department{Name: "Sales"}
	john, carlos := model.Employee{Name: "John"}, model.Employee{Name: "Carlos"}

	// begin trx
	tx := db.MustBegin()

	// insert dept
	res, err := tx.NamedExec(model.NAMEDQUERY_INSERT_DEPT, sales)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	salesDepId, err := res.LastInsertId()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	sales.Id = salesDepId

	// insert empl 1
	john.DepId = salesDepId
	res, err = tx.NamedExec(model.NAMEDQUERY_INSERT_EMPL_DEPTID, john)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	johnId, err := res.LastInsertId()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	john.Id = johnId

	// insert empl 2
	carlos.DepId = salesDepId
	res, err = tx.NamedExec(model.NAMEDQUERY_INSERT_EMPL_DEPTID, carlos)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	carlosId, err := res.LastInsertId()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	carlos.Id = carlosId

	// commit trx
	if err = tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return nil
}

func txDemoAssign(db *sqlx.DB) error {
	// define initial
	finance, biz := model.Department{Name: "Finance"}, model.Department{Name: "Business"}
	denis, gerry := model.Employee{Name: "Denis"}, model.Employee{Name: "Gerry"} // initial: denis -> finance, gerry -> biz

	// insert data without trx
	financeId, err := model.InsertDepartmentAndReturnId(db, finance)
	if err != nil {
		return err
	}
	finance.Id = financeId
	denis.DepId = financeId
	denisId, err := model.InsertEmployeeAndReturnId(db, denis)
	if err != nil {
		return err
	}
	denis.Id = denisId

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

	// insert dept with trx
	tx := db.MustBegin()
	res, err := tx.NamedExec(model.NAMEDQUERY_INSERT_DEPT, biz)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	bizId, err := res.LastInsertId()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	biz.Id = bizId

	// insert employee with trx
	gerry.DepId = bizId
	res, err = tx.NamedExec(model.NAMEDQUERY_INSERT_EMPL_DEPTID, gerry)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	gerryId, err := res.LastInsertId()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	gerry.Id = gerryId

	// update employee with trx
	denis.DepId = bizId
	_, err = tx.NamedExec("update employee set department_id = :department_id where id = :id", denis)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	// commit trx
	if err = tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

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
