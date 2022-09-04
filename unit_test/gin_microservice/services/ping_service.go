package services

import "fmt"

type pingServiceInterface interface {
	HandlePing() (string, error)
}

type pingService struct{}

const (
	pong = "pong"
)

var (
	// public handler / private struct
	PingService pingServiceInterface = pingService{}
)

func (p pingService) HandlePing() (string, error) {
	fmt.Println("Getting from actual service")
	return pong, nil
}
