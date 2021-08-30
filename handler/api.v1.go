package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	. "kovercheng/middleware"
	"kovercheng/service/database"
	"net/http"
)

var call = resty.New().R()

// TODO Error handling
func example(c *gin.Context) {
	Logger.Debug("API been called.")
	// HTTP call via go-resty
	if _, err := call.Get("https://www.microsoft.com/"); err != nil {
		Logger.Warnf("%+v", err)
	}

	if err := database.TestPostgres(); err != nil {
		Logger.Warnf("%+v", err)
	}

	if err := database.TestMongo(); err != nil {
		Logger.Warnf("%+v", err)
	}

	if err := database.TestRedis(); err != nil {
		Logger.Warnf("%+v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "example",
	})
}
