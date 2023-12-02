package router

import (
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.GET("/", Greet)

	v1 := router.Group("/v1")
	v1.GET("/singles", GetPossibleSingles)
	v1.POST("/singles", PostSingle)
	v1.POST("/reset", ResetAllSingles)
	v1.DELETE("/singles", DeleteSingle)

	return router
}
