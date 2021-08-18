package main

import "fmt"

type car struct {
	Brand string
	Year  int
}

func (car car) getBrand() string {
	return car.Brand
}

func mainStructs() {
	// Instanciar un Struct #1
	myCar := car{Brand: "Subaru", Year: 2017}
	fmt.Println(myCar)

	// Instanciar un Struct #1
	var myCar2 car
	myCar2.Brand = "Mazda"
	fmt.Println(myCar2)

	//Llamado funcion getBrand
	fmt.Println(myCar2.getBrand())

}
