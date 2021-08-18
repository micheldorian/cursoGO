package main

import "fmt"

func suma() {

	x := 10
	y := 50

	// Suma
	result := x + y

	fmt.Println("La suma es", result)

	// Resta
	result = y - x

	fmt.Println("La resta es", result)

	// Multiplicacion
	result = y * x

	fmt.Println("La multiplicacion es", result)

	// Division
	result = y / x

	fmt.Println("La division es", result)

	// Modulo, el residuo de la division
	result = y % x

	fmt.Println("El residuo es", result)

	// Incremental
	x++

	fmt.Println("El incremental de x es", x)

	// Decremental
	x--

	fmt.Println("El decremental de x es", x)

}

func resta() {

	// Declaracion de variables enteras
	base := 15
	altura := 10
	var area int
	// Declaracion de variables enteras
	area = base * altura
	fmt.Println("El area del rectangulo es", area)

}

func division() {

	// Declaracion de variables enteras
	base := 15
	altura := 10
	var area int
	// Declaracion de variables enteras
	area = base * altura
	fmt.Println("El area del rectangulo es", area)

}

func multiplicacion() {

	// Declaracion de variables enteras
	base := 15
	altura := 10
	var area int
	// Declaracion de variables enteras
	area = base * altura
	fmt.Println("El area del rectangulo es", area)

}
