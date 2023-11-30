package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}
