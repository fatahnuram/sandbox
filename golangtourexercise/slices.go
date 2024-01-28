package golangtourexercise

// solves exercise at https://go.dev/tour/moretypes/18

// Pic, call it using: `pic.Show(Pic)`
// Import `pic` from `golang.org/x/tour/pic`
func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		arr[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			arr[i][j] = uint8(i ^ j)
			// you may change this function to other function
			// e.g. i*j, i^j, or (i+j)/2
		}
	}
	return arr
}
