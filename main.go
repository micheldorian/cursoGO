package main

import (
	"curso_golang_platzi/restservice"
	"fmt"
)

func main() {
	response := restservice.RestCallGET()
	fmt.Println(response)
}
