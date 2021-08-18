package main

import "fmt"

func sayChannel(text string, c chan<- string) {
	c <- text // Entra el texto al canal

}

func chanel() {
	c := make(chan string, 1) // 1 es la cantidad limite
	fmt.Println("Hello")
	go sayChannel("bye", c)
	fmt.Println(<-c)
}
