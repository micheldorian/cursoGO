package src

import "fmt"

func mainAreas() {
	fmt.Printf("%v es el Area de circulo\n", areaCirculo(2))
	fmt.Printf("%v es el Area de rectangulo\n", areaRectangulo(20, 10))
	fmt.Printf("%v es el Area de triangulo\n", areaTriangulo(20, 10))
	fmt.Printf("%v es el Area de trapezoide\n", areaTrapezoide(18, 22, 10))
	_, valor2 := dobleValor(18)
	fmt.Println("Retornar valor y doble valor", valor2)
}

func areaCirculo(radio float64) float64 {
	// Declaracion constante
	const pi float64 = 3.141516
	return pi * (radio * radio)
}

func areaRectangulo(base, altura int) int {
	return base * altura
}

func areaTriangulo(base, altura int) int {
	return (base * altura) / 2
}

func areaTrapezoide(base1, base2, altura float64) float64 {
	return altura * (base1 + base2) / 2
}

func dobleValor(a int) (b, c int) {
	return a, a * 2
}
