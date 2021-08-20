package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Breed    string
	WeightKg *int
}

func main() {

	d := Dog{Breed: "Pug"}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
	fmt.Println(d.WeightKg)
}
