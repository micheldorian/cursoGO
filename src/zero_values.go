package src

import "fmt"

func zeroValues() {
	// Zero Values
	var a int     // Imprime 0
	var b float64 // Imprime 0
	var c string  // Imprime un string vacio NO es un nulo
	var d bool    // Imprime false

	fmt.Println("var a int  -->", a)
	fmt.Println("var b float64  -->", b)
	fmt.Println("var c string  -->", c, "Imprime un string vacio NO es un nulo")
	fmt.Println("var d bool  -->", d)

}
