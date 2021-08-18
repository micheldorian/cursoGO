package restservice

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type responseStruct struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TestRestCallGET(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		responseMock := &responseStruct{
			UserID:    1,
			ID:        1,
			Title:     "delectus aut autem",
			Completed: false}

		responseData := RestCallGET()
		responseRest := responseStruct{}
		json.Unmarshal(responseData, &responseRest)
		assert.Equal(t, responseMock.Completed, responseRest.Completed)
	})
}
