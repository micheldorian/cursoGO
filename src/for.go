package src

import "fmt"

func forCondition() {
	// Condicional For
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// For while
	counter := 100
	for counter > 95 {
		fmt.Println(counter)
		counter--
	}

	// for range posicion y dato
	arrayPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for posicion, dato := range arrayPares {
		fmt.Println("La posicion", posicion)
		fmt.Println("El numero", dato)
	}

	// for range suma valores
	arrayImpares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	sum := 0
	for _, dato := range arrayImpares {
		sum += dato
	}

	fmt.Println("La suma del array impar es", sum)

	// for range creacion array de impares a partir de array
	arrayNumeros := []int{21, 3, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21}
	var s []int
	for _, dato := range arrayNumeros {
		if dato%2 != 0 {
			s = append(s, dato)
		}
	}
	fmt.Println(s)
}
