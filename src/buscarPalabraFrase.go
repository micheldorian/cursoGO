package src

import (
	"strings"
)

func Coincidencias(frase string, palabra string) int {
	clearText := strings.ToLower(strings.Replace(frase, ",", "", -1))
	cantidad := 0
	if strings.Contains(clearText, palabra) {
		palabras := strings.Split(clearText, " ")
		//fmt.Printf("%T", palabras)
		for _, v := range palabras {
			if v == palabra {
				cantidad++
			}
		}
	}
	return cantidad
}
