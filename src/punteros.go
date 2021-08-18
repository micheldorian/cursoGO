package src

import (
	"fmt"
)

type pcType struct {
	ram   int
	disk  int
	brand string
}

func (pc pcType) getBrand() {
	fmt.Println(pc.brand)
}

func (pc *pcType) duplicateRam() {
	pc.ram = pc.ram * 2
}

func (pc pcType) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de disco y es una %s,", pc.ram, pc.disk, pc.brand)
}

func punteros() {
	// Punteros, llevar funciones de una libreria a otra de manera mas eficiente
	a := 50
	b := &a //Puntero de direccion de memoria de a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // Al valor de la direccion de memoria
	*b = 100        // actualiza a y *b
	fmt.Println(a)

	myPc := pcType{ram: 16, disk: 512, brand: "MSI"}
	fmt.Println(myPc)
	myPc.getBrand()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
