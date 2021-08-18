package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("AddOKResponse", func(t *testing.T) {

		got := Add(2, 1)
		want := 3
		assert.Equal(t, got, want)
	})

	t.Run("Error", func(t *testing.T) {

		got := Add(2, 1)
		want := 10
		assert.NotEqual(t, got, want)
	})
}
