package main

import "fmt"

func applyIt(afunct func(int, int) int, value1 int, value2 int) int {
	return afunct(value1, value2)
}

func main() {
	var numUser int
	var multiplicador int

	fmt.Print("Please, enter a number 1: ")
	fmt.Scan(&numUser)
	fmt.Print("Please, enter a number 2: ")
	fmt.Scan(&multiplicador)

	v := applyIt(func(x, y int) int { return x * y }, numUser, multiplicador)
	fmt.Println("Result of multiply number 1 and number 2 is: :", v)
}
