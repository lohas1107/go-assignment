package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tinder/internal/matching"
)

const (
	QueryKeyMostPossible = "most_possible"
)

func Greet(context *gin.Context) {
	context.JSON(http.StatusOK, "hello, world")
}

func DeleteAllSingles(context *gin.Context) {
	matching.Initialize() //todo
	context.JSON(http.StatusOK, nil)
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
	var single matching.Single
	err := context.BindJSON(&single)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	if !single.IsValidGender() {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	if single.Height <= 0 {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	possibleMatches := matching.AddAndMatch(single)
	context.JSON(http.StatusCreated, possibleMatches)
}
