package main

import "fmt"

func main() {

	// Declaracion de Constantes
	const pi float64 = 3.14
	const pi2 = 3.141516
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int     // Imprime 0
	var b float64 // Imprime 0
	var c string  // Imprime un string vacio NO es un nulo
	var d bool    // Imprime false

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const basecuadrado = 10
	areacuadrado := basecuadrado * basecuadrado

	fmt.Println("El area del cuadrado", areacuadrado)

}
