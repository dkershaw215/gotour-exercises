package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	lastZ := 0.0
	for {
		if math.Abs(z - lastZ) < 0.0000001 {
			return z
		}
		lastZ = z
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
		
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
