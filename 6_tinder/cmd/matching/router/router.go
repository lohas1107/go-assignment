package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.GET("/", Greet)
	v1 := router.Group("/v1")
	v1.GET("/singles")
	return router
}

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}
