package service

import (
	"errors"
	"sync"
)

type service struct{}

var (
	primeCache map[uint32]uint32
	mutex      sync.Mutex
)

type PrimeService interface {
	Calculate(primeNumber uint32) (uint32, error)
}

// Info: Here we can inject database instance as a parameter but thats not required.
func NewPrimeService(dataMap map[uint32]uint32) PrimeService {
	primeCache = dataMap
	return &service{}
}

func (*service) Calculate(primeNumber uint32) (result uint32, err error) {

	if primeNumber == 0 {
		err := errors.New("Zero is invalid input. Please provide valid input.")
		return 0, err
	}

	// Safe to read concurrently. Reading from local map acting as cache object
	if value, found := primeCache[primeNumber]; found {
		return value, nil
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
	writeToLocalCache(givenNumber, result)
	return result
}

func writeToLocalCache(givenPrime uint32, resultantPrime uint32) error {
	mutex.Lock()
	defer mutex.Unlock()

	primeCache[givenPrime] = resultantPrime
	return nil
}
