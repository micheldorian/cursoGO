package src

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func channel2() {
	c := make(chan string, 2)
	c <- "Msj 1"
	c <- "Msj 2"
	fmt.Println(len(c), cap(c)) // len Cuantas GoRoutines hay en el channel, cap cual es la cantidad maximas que puede almacenar el channel.
	//range y close
	close(c)                 // le dice al rountime de go que va a cerrar el canal sin importar si tiene o no mas capacidad
	for message := range c { // nos ayuda a hacer el recorrido por el channel
		fmt.Println(message)
	}
	//Select
	//cuando manejamos multiples canales y no tenemos certeza de cual de los canales responde primero hay que manejar el select.
	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de: ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de: ", m2)
		}
	}
}
