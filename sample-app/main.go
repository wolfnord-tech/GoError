package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for i := 0; i < 1000; i++ {
		prev := z
		z = z - (z*z-x)/(2*z)
		if abs(z-prev) < 1e-10 {
			break
		}
	}
	return z, nil
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
