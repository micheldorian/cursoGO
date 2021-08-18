package main

import "fmt"

func functionFmt() {
	// Declaracion Variables
	nombre := "Platzi"
	cursos := 500

	// Println
	fmt.Println(nombre, cursos)
	// Printf
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)
	// Sprint
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

}
