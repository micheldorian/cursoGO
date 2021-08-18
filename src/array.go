package main

import "fmt"

func array() {

	// for range posicion y dato
	arrayNumeros := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for posicion, dato := range arrayNumeros {
		fmt.Println("La posicion", posicion)
		fmt.Println("El numero", dato)
	}

	// for range suma valores
	arraySuma := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	sum := 0
	for _, dato := range arraySuma {
		sum += dato
	}

	fmt.Println("La suma del array impar es", sum)

	// for range creacion array de impares a partir de array
	arrayImpares := []int{21, 3, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21}
	var s []int
	for _, dato := range arrayImpares {
		if dato%2 != 0 {
			s = append(s, dato)
		}
	}
	fmt.Println(s)
}
