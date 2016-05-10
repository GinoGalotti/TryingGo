package main

import (
	"fmt"
	"math"
)

var treshold = 0.01

func Sqrt(x float64) (z float64) {

	z = float64(1)
	close_enough := false
	for !close_enough {
		
		z_new := z - ((z*z - x) / 2*z)
		
		fmt.Println("Abs = ", math.Abs(z_new - z))
		if math.Abs(z_new - z) < treshold {
		
			close_enough = true
		}
		z = z_new
	}
	return 
}

func main() {
	fmt.Println("FINAL NUMBER: ", Sqrt(4))
}

