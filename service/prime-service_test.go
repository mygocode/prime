package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPrimeCache                = map[uint64]uint64{0: 0}
	testPrimeService PrimeService = NewPrimeService(testPrimeCache)
)

func TestCalculate_WithZero(t *testing.T) {
	_, err := testPrimeService.Calculate(0)

	assert.NotNil(t, err)
	assert.Equal(t, "Zero is invalid input. Please provide valid input.", err.Error())
}

func TestCalculate_TenMillion(t *testing.T) {
	result, err := testPrimeService.Calculate(10000000)
	expected := uint64(9999991)

	assert.Nil(t, err)
	assert.Equal(t, result, expected, "Should be same")
}

func TestCalculate_FiveHundredMillion(t *testing.T) {
	result, err := testPrimeService.Calculate(500000000)
	expected := uint64(499999993)

	assert.Nil(t, err)
	assert.Equal(t, result, expected, "Should be same")
}

func TestCalculate_FiveHundredBillion(t *testing.T) {
	result, err := testPrimeService.Calculate(500000000000)
	expected := uint64(499999999979)

	assert.Nil(t, err)
	assert.Equal(t, result, expected, "Should be same")
}
