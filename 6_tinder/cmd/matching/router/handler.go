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

func GetPossibleSingles(context *gin.Context) {
	count, err := strconv.Atoi(context.Query(QueryKeyMostPossible))
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	if count <= 0 {
		context.JSON(http.StatusOK, []any{})
		return
	}

	possibleMatches := matching.GetPossibleMatches(count)
	context.JSON(http.StatusOK, possibleMatches)
}

func PostSingle(context *gin.Context) {
	var single *matching.Single
	err := context.BindJSON(&single)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	if !single.IsValidGender() || single.Height <= 0 || single.WantedDates <= 0 {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	possibleMatches := matching.AddAndMatch(single)
	context.JSON(http.StatusCreated, possibleMatches)
}

func ResetAllSingles(context *gin.Context) {
	matching.Initialize()
	context.JSON(http.StatusOK, nil)
}

func DeleteSingle(context *gin.Context) {
	// todo
	context.JSON(http.StatusOK, nil)
}
