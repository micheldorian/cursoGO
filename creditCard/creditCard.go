package creditcard

import "fmt"

type CreditCard struct {
	Brand  string `json:"brand,omitempty"`
	Number string `json:"number,omitempty"`
	CVV    string `json:"cvv,omitempty"`
	Expiry string `json:"expiry,omitempty"`
	Status string `json:"status,omitempty"`
}

type DebitCard struct {
	Brand  string `json:"brand,omitempty"`
	Number string `json:"number,omitempty"`
	Expiry string `json:"expiry,omitempty"`
	Status string `json:"status,omitempty"`
}

type products interface {
	GetActivate() string
	GetNumber() string
}

func (d DebitCard) GetActivate() string {
	return d.Status
}

func (c CreditCard) GetActivate() string {
	return c.Status
}

func (d DebitCard) GetNumber() string {
	return d.Number
}

func (c CreditCard) GetNumber() string {
	return c.Number
}

func NewCreditCard(brand, number, expiry, cvv string) *CreditCard {
	card := &CreditCard{
		Brand:  brand,
		Number: number,
		CVV:    cvv,
		Expiry: expiry,
	}
	return card
}

func NewDebitCard(brand, number, expiry string) *DebitCard {
	card := &DebitCard{
		Brand:  brand,
		Number: number,
		Expiry: expiry,
	}
	return card
}

func NewCreditCardValidate(brand, number, expiry, cvv string) (*CreditCard, error) {
	card := &CreditCard{
		Brand:  brand,
		Number: number,
		CVV:    cvv,
		Expiry: expiry,
	}
	if err := card.Validate(); err != nil {
		return nil, err
	}
	return card, nil
}

func GetActivateProduct(p products) {
	fmt.Println("El estado del producto es: ", p.GetActivate())
}

func GetProductNumber(p products) {
	fmt.Println("El numero del producto es: ", p.GetNumber())
}
func ProductsFunc() {

	card := NewCreditCard("VISA", "5345687904939234", "01/24", "432")
	debit := NewDebitCard("MASTER", "4345687904939234", "01/25")

	GetActivateProduct(card)
	GetActivateProduct(debit)
	GetProductNumber(card)
	GetProductNumber(debit)
}
