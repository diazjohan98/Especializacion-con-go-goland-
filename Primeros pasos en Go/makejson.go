package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter a name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Enter an address: ")
	scanner.Scan()
	address := scanner.Text()

	datos := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(datos)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

fmt.Println(string(jsonData))
}