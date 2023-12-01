package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	QueryKeyMostPossible = "most_possible"
)

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}

func GetPossibleSingles(context *gin.Context) {
	mostPossible, err := strconv.Atoi(context.Query(QueryKeyMostPossible))
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	if mostPossible <= 0 {
		context.JSON(http.StatusOK, []any{})
	}
}

func PostSingle(context *gin.Context) {
	context.JSON(http.StatusCreated, []any{})
}
