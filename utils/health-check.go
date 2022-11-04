package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health Check
// @Accept */*
// @Produce json
// @Success 200 {string} Server is up and running
// @Router /health-check [GET]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"message": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
