package golangtourexercise

// solves exercise at https://go.dev/tour/moretypes/26

// Fibonacci is a function that returns
// a function that returns an int.
//
// call it using: ```
// f := Fibonacci()
//
//	for i := 0; i < 15; i++ {
//		fmt.Println(f())
//	}
//
// ```
func Fibonacci() func() int {
	counter := 0
	latest1 := 1
	latest2 := 0
	sum := 0

	return func() int {
		switch counter {
		case 0:
			counter++
			return 0
		case 1:
			counter++
			return 1
		default:
			counter++
			sum = latest2 + latest1
			latest2 = latest1
			latest1 = sum
			return sum
		}
	}
}
