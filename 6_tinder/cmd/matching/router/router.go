package router

import (
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.GET("/", Greet)

	v1 := router.Group("/v1")
	v1.GET("/singles", GetSingles)

	return router
}
