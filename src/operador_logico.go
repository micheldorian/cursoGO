package src

import "fmt"

func operadorLogicoIf() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 != valor2 {
		fmt.Println("Los valores son diferentes")
	}
	//With AND
	if valor1 != 2 && valor2 == 2 {
		fmt.Println("Caso AND es verdadero")
	}
	//With OR
	if valor1 == 3 || valor2 == 2 {
		fmt.Println("Caso OR al menos uno es verdadero")
	}
}
func operadorLogicoSwitch() {
	dato := 5 % 2
	switch dato {
	case 0:
		fmt.Println("Es pas")
	default:
		fmt.Println("Es impar")
	}

	// Con Condicion
	switch dato := 5 % 2; dato {
	case 0:
		fmt.Println("Es pas")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	dato2 := 99
	switch {
	case dato2 < 100:
		fmt.Println("Es Menor a 100")
	case dato2 > 100:
		fmt.Println("Es mayor a 100")
	default:
		fmt.Println("Es 100")
	}
}
