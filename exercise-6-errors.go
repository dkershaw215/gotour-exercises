package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error)  {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := x / 2
	lastZ := 0.0
	for {
		if math.Abs(z - lastZ) < 0.0000001 {
			return z, nil
		}
		lastZ = z
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
		
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-4))
}
