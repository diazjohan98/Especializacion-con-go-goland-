package main

import "fmt"

func getMax(vale ...int) int {
	maxV := -1
	for _, v := range vale {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	fmt.Println(getMax(1, 3, 6, 7))
	vslices := []int{1, 3, 6, 7}
	fmt.Println(getMax(vslices...))
}
