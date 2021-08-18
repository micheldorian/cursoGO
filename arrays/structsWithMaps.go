package arrays

import (
	def "curso_golang_platzi/definitions"
	"fmt"
)

var cars = map[string]def.CarType{}

func addCar(car def.CarType, placa string) {
	cars[placa] = car
	fmt.Println("Car add")
}

func getCar(placa string) def.CarType {
	value := cars[placa]
	return value
}
func deleteCar(placa string) {
	delete(cars, placa)
	fmt.Println("Car delete")
}

func getCars() map[string]def.CarType {
	return cars
}

func StructsWithMaps() {
	//Map de Struct
	myCar := def.CarType{Brand: "Subaru", Year: 2017}
	myCar2 := def.CarType{Brand: "Mazda", Year: 2020}

	def.UpdateCar(&myCar, "Subaru XV", 2018)

	addCar(myCar, "YUI-123")
	addCar(myCar2, "AAS-433")

	fmt.Println(getCars())

	fmt.Println(getCar("YUI-123"))

	fmt.Println(getCar("AAS-433"))

	deleteCar("AAS-433")

	fmt.Println(getCars())

	fmt.Println(getCar("AAS-431"))

	fmt.Println(myCar.GetBrand())

	fmt.Println(myCar.GetYear())

}
