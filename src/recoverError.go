package main

import (
	"fmt"
	"strconv"
)

func cadenaANumero(cadena string) int {
	defer func() {
		recuperado := recover()
		if recuperado != nil {
			fmt.Println("Proceso recover, manejo de error")
		}
	}()
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		panic(err)
	}
	return numero
}

func errorRecover() {
	fmt.Println("Inicia")           // Inicia
	fmt.Println(cadenaANumero("2")) // 2
	fmt.Println(cadenaANumero("j")) // 0
	fmt.Println("Fin")              // Fin
}
