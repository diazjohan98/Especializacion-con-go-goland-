package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	slice1 := make([]int, 3)
	
	var inp string
	var count int
	var exit_cond bool = true
	count = 0
	for exit_cond {
		fmt.Print("Please enter number: ")
		fmt.Scanln(&inp)

		if inp == "X" {
			exit_cond = false
		} else {

			num, _ := strconv.Atoi(inp)

			
			slice1 = append(slice1, num)
			

			sort.Sort(sort.IntSlice(slice1))
			
			fmt.Println("Sorted list will be: ")
			fmt.Println(slice1)

			count++
		}
	}

	
	

	
	

}

