package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}

func GetSingles(context *gin.Context) {
	context.JSON(http.StatusBadRequest, nil)
}
