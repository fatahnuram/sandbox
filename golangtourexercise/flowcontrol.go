package golangtourexercise

// solves exercise at https://go.dev/tour/flowcontrol/8

import (
	"fmt"
)

// Sqrt, call it using: `fmt.Println(Sqrt(1024))`
func Sqrt(x float64) float64 {
	res := 0.0
	z := 1.0
	awal := 0.0
	akhir := 1.0
	i := 0
	for awal != akhir {
		awal = z
		fmt.Println("awal", awal)
		z -= (z*z - x) / (2 * z)
		akhir = z
		fmt.Println("akhir", akhir)
		res = z
		i++
	}
	fmt.Println(i)
	return res
}
