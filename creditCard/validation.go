package creditcard

import (
	"fmt"
	"strconv"
	"time"
)

// Validate implements various checks to ensure a credit card is valid.
func (card *CreditCard) Validate() error {
	const shortDate = "01/2006"
	t, err := time.Parse(shortDate, card.Expiry)
	if err != nil {
		return fmt.Errorf("Credit Card date invalid %v", t)
	}
	// Validate: card expired
	if t.Year() < time.Now().Year() {
		return fmt.Errorf("Credit Card expired: %v", card.Expiry)
	}
	// Validate: parsing to integer for number and cvv2
	if _, err := strconv.Atoi(card.Number); err != nil {
		return fmt.Errorf("Credit Card number invalid: %v", card.Number)
	}
	if _, err := strconv.Atoi(card.CVV); err != nil {
		return fmt.Errorf("Credit Card CVV2 invalid: %v", card.CVV)
	}
	// Validate: cvv2
	if len(card.CVV) < 3 || len(card.CVV) > 4 {
		return fmt.Errorf("Credit Card CVV2 length invalid: %v", card.CVV)
	}
	// Validate: card length
	if len(card.Number) < 13 || len(card.Number) > 19 {
		return fmt.Errorf("Credit Card number length invalid: %v", card.Number)
	}
	return nil
}
