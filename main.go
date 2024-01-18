package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	/* contoh Println vs Sprintln */

	// Println
	fmt.Println("Example of Println function.")

	// Sprintln (tapi tidak di-print)
	fmt.Sprintln("Example of Sprintln function.")
	// Sprintln (dan di-print)
	s := fmt.Sprintln("Another example of Sprintln function.")
	io.WriteString(os.Stdout, s)

	/* contoh fmt.Errorf vs errors.New */

	// fmt.Errorf
	errcontent := "contoh dynamic error msg"
	err := fmt.Errorf("example Errorf msg: %s", errcontent)
	fmt.Println(err)

	// errors.New
	err2 := errors.New("example errors.New for static error msg")
	fmt.Println(err2)
}
