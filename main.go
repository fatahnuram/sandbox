package main

import (
	"log"
	"os"

	"github.com/fatahnuram/sandbox/db"
)

func main() {
	_, err := db.InitializeMysql("127.0.0.1:3306", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), "tcp")
	if err != nil {
		log.Fatalf("failed to init database, err: %v", err)
	}
	log.Println("connected to db.")
}
