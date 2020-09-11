package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Note:
If we had data layer, we would have tested service functions with mocked database objects.
This will help us to test service functions with database dependency.
*/

func TestCalculate(t *testing.T) {
	testPrimeService := NewPrimeService()
	_, err := testPrimeService.Calculate(0)

	assert.NotNil(t, err)
	assert.Equal(t, "Zero is invalid input. Please provide valid input.", err.Error())
}
