package main

import (
	"unit_test/gin_microservice/controller"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controller.Ping)
	router.Run(":8080")
}
