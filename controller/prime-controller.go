package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mygocode/prime/errors"
	"github.com/mygocode/prime/service"
)

type controller struct{}

var (
	primeService service.PrimeService
)

type PrimeController interface {
	PostPrime(response http.ResponseWriter, request *http.Request)
}

func NewPrimeController(service service.PrimeService) PrimeController {
	primeService = service
	return &controller{}
}

func (*controller) PostPrime(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	receivedPrime, err1 := strconv.ParseInt(request.PostFormValue("primenumber"), 0, 32)
	if err1 != nil {
		fmt.Println(err1)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Invalid input. Please provide valid numeric value."})
		return
	}

	result, err2 := primeService.Calculate(uint32(receivedPrime))
	if err2 != nil {
		fmt.Println(err2)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error processing the given information"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
