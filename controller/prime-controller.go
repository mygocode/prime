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

	receivedPrime, err1 := strconv.ParseUint(request.PostFormValue("primenumber"), 10, 64)
	if err1 != nil {
		fmt.Println("Controller->" + err1.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Invalid input. Please provide valid numeric value."})
		return
	}

	result, err2 := primeService.Calculate(uint32(receivedPrime))
	if err2 != nil {
		fmt.Println("Controller-> " + err2.Error())
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
