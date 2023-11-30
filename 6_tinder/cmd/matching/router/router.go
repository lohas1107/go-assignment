package router

import "github.com/gin-gonic/gin"

func SetUp() *gin.Engine {
	router := gin.Default()
	return router
}
