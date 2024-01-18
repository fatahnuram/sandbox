package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	// contoh Println vs Sprintln
	// Println
	fmt.Println("Example of Println function.")

	// Sprintln (tapi tidak di-print)
	fmt.Sprintln("Example of Sprintln function.")
	// Sprintln (dan di-print)
	s := fmt.Sprintln("Another example of Sprintln function.")
	io.WriteString(os.Stdout, s)
}
