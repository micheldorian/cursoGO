package src

import (
	"fmt"
	"strings"
)

func mainPalindromo() {
	isPalindromo("Ama")
}

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}
