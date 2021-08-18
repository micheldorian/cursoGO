package main

import "fmt"

type Cuadrado struct {
	base float64
}

type Rectangulo struct {
	base   float64
	altura float64
}

type figuras2D interface {
	area() float64
}

func (c Cuadrado) area() float64 {
	return c.base * c.base
}

func (r Rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("El area es: ", f.area())
}

func interfaces() {
	myCuadrado := Cuadrado{base: 2}
	myRectangulo := Rectangulo{base: 2, altura: 4}
	calcular(myCuadrado)
	calcular(myRectangulo)

	//lISTA DE INTERFACES
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
