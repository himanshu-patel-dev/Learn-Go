package controller

import (
	"net/http"
	"unit_test/gin_microservice/services"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, result)
}
