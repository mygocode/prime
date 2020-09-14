package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPrimeCache                = map[uint32]uint32{0: 0}
	testPrimeService PrimeService = NewPrimeService(testPrimeCache)
)

func TestCalculate_WithZero(t *testing.T) {
	_, err := testPrimeService.Calculate(0)

	assert.NotNil(t, err)
	assert.Equal(t, "Zero is invalid input. Please provide valid input.", err.Error())
}

func TestCalculate_TenMillion(t *testing.T) {
	result, err := testPrimeService.Calculate(10000000)
	expected := uint32(9999991)

	assert.Nil(t, err)
	assert.Equal(t, result, expected, "Should be same")
}

func TestCalculate_FiveHundredMillion(t *testing.T) {
	result, err := testPrimeService.Calculate(500000000)
	expected := uint32(499999993)

	assert.Nil(t, err)
	assert.Equal(t, result, expected, "Should be same")
}
