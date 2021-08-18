package src

import "fmt"

func keywords() {
	//defer, imprime esa linea de ultimo antes de que el proceso termine
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//continue cuando quiero que algo continue a pesar de que pase un error
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//break cuando quiero que algo termine o no se ejecute a pesar de que pase un error
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
