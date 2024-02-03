package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
	}

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db.")

	fmt.Println("\n======= Demo Insert")
	// insert to department
	dep := Department{Name: "IT"}
	depId, err := insertDepartmentAndReturnId(db, dep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New department ID: %d\n", depId)

	// insert to employee
	empl := Employee{Name: "Putra", DepId: depId}
	emplId, err := insertEmployeeAndReturnId(db, empl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New employee ID: %d\n", emplId)

	fmt.Println("\n======= Demo Retrieve")
	// get department
	depIT, err := getDepartmentById(db, depId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Department with ID %d: %v\n", depIT.Id, depIT)

	// get employee
	myEmployee, err := getEmployeeById(db, emplId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee with ID %d: %v\n", myEmployee.Id, myEmployee)

	fmt.Println("\n======= Demo Update")
	// update department
	updateDepartmentNameById(db, depId, "Tech")
	depIT, err = getDepartmentById(db, depId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated department, ID: %d, obj: %v\n", depIT.Id, depIT)

	// update employee
	updateEmployeeNameById(db, emplId, "Kamaludin")
	myEmployee, err = getEmployeeById(db, emplId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated employee, ID: %d, obj: %v\n", myEmployee.Id, myEmployee)

	fmt.Println("\n======= Demo Update")
	// delete employee
	deletedEmployee, err := deleteAndReturnEmployeeById(db, emplId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted employee: %v\n", deletedEmployee)

	// delete department
	deletedDept, err := deleteAndReturnDepartmentById(db, depId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted department: %v\n", deletedDept)

	// tx 1 dept 2 empl
	fmt.Println("\n======= Demo Transaction 1 Dept 2 Empl")
	txOneDeptTwoEmpl(db)
	printAllDept(db)
	printAllEmpl(db)

	// tx demo assign empl to new dept
	fmt.Println("\n======= Demo Transaction Assign Empl to new Dept")
	txDemoAssign(db)
}
