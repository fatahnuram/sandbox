package golangtourexercise

// solves exercise at https://go.dev/tour/moretypes/23

import (
	"strings"
)

// WordCount, call it using: `wc.Test(WordCount)`
// Import `wc` from `golang.org/x/tour/wc`
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)
	for _, w := range arr {
		v, ispresent := m[w]
		if !ispresent {
			m[w] = 1
		} else {
			m[w] = v + 1
		}
	}
	return m
}
