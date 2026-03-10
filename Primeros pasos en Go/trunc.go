package main

import (
	"fmt"
)

func main() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error al leer la entrada.")
		return
	}

	if len(input) >= 3 {
		
		primeraLetra := input[0]
		ultimaLetra := input[len(input)-1]

		empiezaI := (primeraLetra == 'i' || primeraLetra == 'I')
		terminaN := (ultimaLetra == 'n' || ultimaLetra == 'N')

		if empiezaI && terminaN {
			tieneA := false
			for _, letra := range input {
				if letra == 'a' || letra == 'A' {
					tieneA = true
					break
				}
			}

			if tieneA {
				fmt.Println("¡Encontrado!")
				return 
			}
		}
	}

	fmt.Println("¡No encontrado!")
}