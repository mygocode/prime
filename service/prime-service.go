package service

type PrimeService interface {
	//Validate(post *entity.Post) error
	//Create(post *entity.Post) (*entity.Post, error)
	// FindAll() ([]entity.Post, error)
	//Calculate(prime *entity.PrimeDomain) (uint32, error)
	Calculate(primeNumber uint32) (uint32, error)
}

type service struct{}

// var (
// 	repo repository.PostRepository
// )

// Info: Here we can inject database instance as a parameter but thats not required.
func NewPrimeService() PrimeService {
	return &service{}
}

// func (*service) Calculate(prime *entity.PrimeDomain) (result uint32, err error) {

// 	return getPrime(prime.PrimeNumber), nil
// }

func (*service) Calculate(primeNumber uint32) (result uint32, err error) {

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
