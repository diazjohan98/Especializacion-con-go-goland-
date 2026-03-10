package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string

	fmt.Printf("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() 

	var names []Name

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		parts := strings.Fields(line)
		
		if len(parts) >= 2 {
			newName := Name{
				fname: parts[0],
				lname: parts[1],
			}
			names = append(names, newName)
		}
	}

	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}