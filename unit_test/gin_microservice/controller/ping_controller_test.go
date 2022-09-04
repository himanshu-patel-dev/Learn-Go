package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"unit_test/gin_microservice/services"

	"github.com/gin-gonic/gin"
)

type mockPingService struct {
	handlePingFn func() (string, error)
}

func (m mockPingService) HandlePing() (string, error) {
	fmt.Println("Inside mocked HandlePing")
	return m.handlePingFn()
}

func TestPingNoError(t *testing.T) {
	services.PingService = mockPingService{handlePingFn: func() (string, error) {
		return "pong", nil
	}}

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}
	if response.Body.String() != "pong" {
		t.Error("response body should be pong")
	}
}

func TestPingWithError(t *testing.T) {
	services.PingService = mockPingService{handlePingFn: func() (string, error) {
		return "", errors.New("Error executing Ping")
	}}

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Ping(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")
	}
	if response.Body.String() != "Error executing Ping" {
		t.Error("response body should be empty")
	}
}
