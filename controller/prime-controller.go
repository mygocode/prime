package controller

import (
	"encoding/json"
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
	//prime := request.FormValue("primenumber")
	receivedPrime, _ := strconv.ParseInt(request.PostFormValue("primenumber"), 0, 32)

	result, err2 := primeService.Calculate(uint32(receivedPrime))
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
