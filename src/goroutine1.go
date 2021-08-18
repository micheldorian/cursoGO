package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) { // Al valor de la direccion de memoria de wg
	defer wg.Done() //defer, imprime esa linea de ultimo antes de que el proceso termine
	fmt.Println(text)
}
func goRoutine1() {
	var wg sync.WaitGroup // Espera que una colección de goroutines se ejecute
	fmt.Println("Hello")
	wg.Add(1)            // Indica que debe esperar a que se ejecute 1 go routine antes de que el main muera
	go say("World", &wg) //Puntero de direccion de memoria de wg
	wg.Wait()            // decirle a la go routine principal del main que espere, hasta que todas las ejecuciones del wg se ejecuten.
}
