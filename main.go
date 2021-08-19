package main

import (
	creditcard "curso_golang_platzi/creditCard"
	"fmt"
	"log"
)

func main() {

	card := creditcard.NewCreditCard("VISA", "5345687904939234", "01/2023", "432")
	debit := creditcard.NewDebitCard("MASTER", "4345687904939234", "01/2030")
	// Validate a CreditCard using all the checks implemented so far.
	if err := card.Validate(); err != nil {
		log.Fatal(err)
	}

	creditcard.GetActivateProduct(card)
	creditcard.GetActivateProduct(debit)
	creditcard.GetProductNumber(card)
	creditcard.GetProductNumber(debit)

	card2, err := creditcard.NewCreditCardValidate("MASTER", "1345687904939234", "01/2022", "112")
	if err != nil {
		log.Fatal(err)
	}

	creditcard.GetActivateProduct(card2)
	creditcard.GetProductNumber(card2)

	encrypt := card2.Encrypt("mySuperSecureSalt")

	// Common helpers.
	fmt.Println(card2.First6())
	fmt.Println(card2.Last4())
	fmt.Println(card2.ToJSON())

	fmt.Println(string(encrypt))
}
