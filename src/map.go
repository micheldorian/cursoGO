package src

import (
	"fmt"
)

func mapFunction() {
	//Mapa o diccionario
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepe"] = 20
	m["Michel"] = 20
	fmt.Println(m)

	//Recorrer un map
	for _, v := range m {
		fmt.Println(v)
	}

	// Encontrar valor
	value, ok := m["Joses"]
	if ok == true {
		fmt.Println("El valor de Jose es ", value)
	} else {
		fmt.Println("No exitste en el map")
	}

	//Update
	m["Jose"] = 40

	//Delete
	delete(m, "Michel")
	fmt.Println(m)
}
