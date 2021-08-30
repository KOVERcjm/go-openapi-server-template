package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"kovercheng/driver"
	"kovercheng/handler"
	. "kovercheng/middleware"
	"os"
)

type option func(*gin.Engine)

var options []option

func include(opts ...option) {
	options = append(options, opts...)
}

func main() {
	router := gin.New()
	router.Use(GinLogger())
	router.Use(gin.Recovery())

	include(handler.V1Router)
	for _, opt := range options {
		opt(router)
	}

	if err := router.Run(os.Getenv("SERVER_URL")); err != nil {
		Logger.Fatalf("Gin start server ERROR: %s", err)
		panic(err)
	}
	_ = driver.CloseConnection()
}
