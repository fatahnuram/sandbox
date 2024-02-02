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

	// TODO: get department
	// TODO: get employee

	// TODO: update department
	// TODO: update employee

	// TODO: delete department
	// TODO: delete employee
}
