package main

import (
	"fmt"
	"math"
)


func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	fmt.Println("Final displacement is:", fn(t))
}