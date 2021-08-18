package restservice

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RestCallGET() (res []byte) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData

}
