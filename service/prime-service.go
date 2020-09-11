package service

import "errors"

type PrimeService interface {
	Calculate(primeNumber uint32) (uint32, error)
}

type service struct{}

// Info: Here we can inject database instance as a parameter but thats not required.
func NewPrimeService() PrimeService {
	return &service{}
}

func (*service) Calculate(primeNumber uint32) (result uint32, err error) {

	if primeNumber == 0 {
		err := errors.New("Zero is invalid input. Please provide valid input.")
		return 0, err
	}

	return getPrime(primeNumber), nil
}

func getPrime(givenNumber uint32) (largestAvailablePrime uint32) {
	var i, j, count, result uint32 = 0, 0, 0, 0

	for i = 1; i < givenNumber; i++ {
		count = 0
		result = givenNumber - i // generating numbers below the input value

		for j = 1; j <= result; j++ {
			if result%j == 0 { //Prime number check
				count++
				if count > 2 {
					break
				}
			}
		}

		if count == 2 {
			break
		}
	}
	return result
}
