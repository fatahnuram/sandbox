package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fatahnuram/sandbox/db"
	selfhttp "github.com/fatahnuram/sandbox/http"
)

func main() {
	// init db connection
	log.Println("initializing db..")
	_, err := db.InitializeMysql("127.0.0.1:3306", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), "tcp")
	if err != nil {
		log.Fatalf("failed to init database, err: %v", err)
	}
	log.Println("connected to db.")

	// start http server
	listen := ":8080"
	log.Printf("listening on %v ..", listen)
	http.ListenAndServe(listen, selfhttp.InitRoutes())
}
