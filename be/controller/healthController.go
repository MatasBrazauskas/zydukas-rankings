package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(g *gin.Context) {
	status := map[string]string{
		"status": "running",
	}

	g.JSON(http.StatusOK, gin.H{"data": status})
}
