package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	router := gin.Default()

	router.GET("/", healthCheck)
	router.GET("/ping", healthCheck)

	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Anchovy API running",
	})
}
