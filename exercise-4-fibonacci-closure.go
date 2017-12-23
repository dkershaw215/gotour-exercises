package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	nPlus := 0
	firstZero := true
	return func() (v int) {
		if firstZero {
			firstZero = false
			v = 0
		} else if n == 0 && nPlus == 0 {
			nPlus = 1
			v = n + nPlus
		} else {
			v = n + nPlus
			n = nPlus
			nPlus = v
		}
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
