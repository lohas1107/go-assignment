package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	"strconv"
	"testing"
	"tinder/cmd/matching/router"
	"tinder/internal/matching"
)

const (
	Host    = "http://localhost:8080"
	Version = "v1"
)

var (
	test *testing.T
)

func GetUrl(path string) string {
	return fmt.Sprintf("%s/%s%s", Host, Version, path)
}

func Reset() apitest.Result {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(GetUrl("/reset")).
		Expect(test).
		Status(http.StatusOK).
		End()
}

func QueryMostPossibleMatches(count int) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(GetUrl("/singles")).
		Query(router.QueryKeyMostPossible, strconv.Itoa(count)).
		Expect(test)
}

func GivenSingleAdded(gender string, height int, wantedDates int) apitest.Result {
	return GivenSingleAdding(gender, height, wantedDates).End()
}

func GivenSingleAdding(gender string, height int, wantedDates int) *apitest.Response {
	single := &matching.Single{
		Name:        "ABC",
		Gender:      gender,
		Height:      height,
		WantedDates: wantedDates,
	}

	request, err := json.Marshal(single)
	if err != nil {
		panic(err)
	}

	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(GetUrl("/singles")).
		Body(string(request)).
		Expect(test)
}

func ShouldResponseBadRequest(response *apitest.Response) apitest.Result {
	return response.
		Status(http.StatusBadRequest).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func ShouldResponseEmptyMatches(response *apitest.Response, status int) apitest.Result {
	return response.
		Status(status).
		Assert(jsonpath.Len("$", 0)).
		End()
}

func AssertMatchesLength(response *apitest.Response, status int, length int) *apitest.Response {
	return response.
		Status(status).
		Assert(jsonpath.Len("$", length))
}

func AssertMatchesContent(assert *apitest.Response, index int, gender string, height int, wantedDates int) *apitest.Response {
	return assert.
		Assert(jsonpath.Equal(fmt.Sprintf("$[%v].gender", index), gender)).
		Assert(jsonpath.Equal(fmt.Sprintf("$[%v].height", index), float64(height))).
		Assert(jsonpath.Equal(fmt.Sprintf("$[%v].wantedDates", index), float64(wantedDates)))
}
