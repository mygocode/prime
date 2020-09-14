package service

import (
	"errors"
	"math/big"
	"sync"
)

type service struct{}

var (
	primeCache map[uint64]uint64
	mutex      sync.Mutex
)

type PrimeService interface {
	Calculate(primeNumber uint64) (uint64, error)
}

// Info: Here we can inject database instance as a parameter but thats not required.
func NewPrimeService(dataMap map[uint64]uint64) PrimeService {
	primeCache = dataMap
	return &service{}
}

func (*service) Calculate(primeNumber uint64) (result uint64, err error) {

	if primeNumber == 0 {
		err := errors.New("Zero is invalid input. Please provide valid input.")
		return 0, err
	}

	// Safe to read concurrently. Reading from local map acting as cache object
	if value, found := primeCache[primeNumber]; found {
		return value, nil
	}

	return getLargestAvailablePrime(primeNumber), nil
}

func getLargestAvailablePrime(givenNumber uint64) (largestAvailablePrime uint64) {

	var result int64 = 1
	num := int64(givenNumber)
	for i := num; i > 1; num-- {

		i := big.NewInt(num)
		isPrime := i.ProbablyPrime(0)

		if isPrime == true {
			result = num
			break
		}
	}

	writeToLocalCache(givenNumber, uint64(result))
	return uint64(result)
}

func writeToLocalCache(givenPrime uint64, resultantPrime uint64) error {
	mutex.Lock()
	defer mutex.Unlock()

	primeCache[givenPrime] = resultantPrime
	return nil
}
