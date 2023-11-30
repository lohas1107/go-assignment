package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.GET("/", Greet)
	router.Group("/v1")
	return router
}

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}
