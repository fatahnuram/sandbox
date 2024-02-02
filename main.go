package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// fmt.Println("Hello, world!")

	// /* Level 1 */
	// /* contoh Println vs Sprintln */

	// // Println
	// fmt.Println("Example of Println function.")

	// // Sprintln (tapi tidak di-print)
	// fmt.Sprintln("Example of Sprintln function.")
	// // Sprintln (dan di-print)
	// s := fmt.Sprintln("Another example of Sprintln function.")
	// io.WriteString(os.Stdout, s)

	// /* contoh fmt.Errorf vs errors.New */

	// // fmt.Errorf
	// errcontent := "contoh dynamic error msg"
	// err := fmt.Errorf("example Errorf msg: %s", errcontent)
	// fmt.Println(err)

	// // errors.New
	// err2 := errors.New("example errors.New for static error msg")
	// fmt.Println(err2)

	// /* Level 2 */
	// generateNIK("akhwat", 2025, 2, 3)
	// generateNIKLanjutan("ARN201-00035", 3)

	/* Tutorial SQL */
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumById(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albId, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albId)
}
