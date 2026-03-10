package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numeros []int
	var entrada string

	fmt.Println("Introduce números enteros uno a uno (escribe 'X' para salir):")

	for {
		_, err := fmt.Scan(&entrada)
		
		if err != nil {
			break
		}

		if entrada == "X" || entrada == "x" {
			fmt.Println("Saliendo del programa...")
			break
		}

		num, err := strconv.Atoi(entrada)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, introduce un número entero.")
			continue 
		}

		numeros = append(numeros, num)

		sort.Ints(numeros)

		fmt.Println(numeros)
	}
}