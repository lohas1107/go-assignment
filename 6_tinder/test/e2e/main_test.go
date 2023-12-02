package e2e

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
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
		Delete(GetUrl("/singles")).
		Expect(test).
		Status(http.StatusOK).
		End()
}

func GivenSingleAdded(gender string, height int, wantedDates int) apitest.Result {
	return GivenSingleAdding(gender, height, wantedDates).End()
}

func GivenSingleAdding(gender string, height int, wantedDates int) *apitest.Response {
	single := &matching.Single{
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
