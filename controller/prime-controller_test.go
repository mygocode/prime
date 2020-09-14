package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mygocode/prime/service"
	"github.com/stretchr/testify/assert"
)

var (
	testPrimeCache                       = map[uint64]uint64{0: 0}
	primeSrv        service.PrimeService = service.NewPrimeService(testPrimeCache)
	primeController PrimeController      = NewPrimeController(primeSrv)
)

func TestPostPrime_BestCase(t *testing.T) {

	reader := strings.NewReader("primenumber=55")
	req, _ := http.NewRequest("POST", "/postprime", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handler := http.HandlerFunc(primeController.PostPrime)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var result uint64
	var expected uint64 = 53
	json.NewDecoder(io.Reader(response.Body)).Decode(&result)

	assert.Equal(t, expected, result)

}

func TestPostPrime_NoNumber(t *testing.T) {

	reader := strings.NewReader("primenumber=")
	req, _ := http.NewRequest("POST", "/postprime", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handler := http.HandlerFunc(primeController.PostPrime)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestPostPrime_String(t *testing.T) {

	reader := strings.NewReader("primenumber=FiveHundred")
	req, _ := http.NewRequest("POST", "/postprime", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handler := http.HandlerFunc(primeController.PostPrime)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)

	status := response.Code
	if status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
