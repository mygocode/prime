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
	testPrimeCache                       = map[uint32]uint32{0: 0}
	primeSrv        service.PrimeService = service.NewPrimeService(testPrimeCache)
	primeController PrimeController      = NewPrimeController(primeSrv)
)

func TestPostPrime_BestCase(t *testing.T) {
	// Create new HTTP request
	reader := strings.NewReader("primenumber=55")
	req, _ := http.NewRequest("POST", "/postprime", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(primeController.PostPrime)

	// Record the HTTP Response
	response := httptest.NewRecorder()

	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var result uint32
	var expected uint32 = 53
	json.NewDecoder(io.Reader(response.Body)).Decode(&result)

	// // Assert HTTP response
	assert.Equal(t, expected, result)

}
