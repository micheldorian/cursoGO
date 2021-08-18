package main

import "fmt"

func listas() {
	//Array
	//Sin tamano definido
	var arrray []int
	fmt.Println(arrray, len(arrray), cap(arrray)) // len cantidad de elementos que tiene arreglo, cap la cantidad maxima de elementos que permite el arreglo
	//Con tamano definido
	var arrray2 [3]int
	fmt.Println(arrray2)
	arrray2[0] = 1
	arrray2[1] = 2
	fmt.Println(arrray2, len(arrray2), cap(arrray2))

	//Slice
	arrayImpares := []int{21, 3, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21}

	println(arrayImpares)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	println(slice)

	//Metodo en el Slice
	fmt.Println(slice[0])   //Imprimier elemento del indice 0
	fmt.Println(slice[:3])  //Imprimir antes del elemento del indice 3. El ultimo se excluye
	fmt.Println(slice[2:4]) //Imprimir el elemento del indice 2 y 3. El ultimo se excluye
	fmt.Println(slice[3:])  //Imprimir desde el elemento del indice 3 hasta el ultimo elemento.

	//Append
	//Agregar un espacio en la lista slice
	slice = append(slice, 7)
	fmt.Println(slice)
	//Agregar slice dentro de slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
	//Recorrer slice
	for _, dato := range slice {
		fmt.Println(dato)
	}
}
