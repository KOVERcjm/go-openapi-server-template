package handler

import (
	"github.com/gin-gonic/gin"
)

func V1Router(e *gin.Engine) {
	router := e.Group("/api/v1")
	{
		router.GET("/example", example)
	}
}
