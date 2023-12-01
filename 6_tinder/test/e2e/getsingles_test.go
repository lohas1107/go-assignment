package e2e

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"tinder/cmd/matching/router"
	"tinder/internal/matching"
)

type GetSinglesTestSuite struct {
	suite.Suite
	Url string
}

func TestGetSinglesTestSuite(t *testing.T) {
	suite.Run(t, new(GetSinglesTestSuite))
	test = t
}

func (s *GetSinglesTestSuite) SetupTest() {
	s.Url = GetUrl("/singles")
	Reset()
}

func (s *GetSinglesTestSuite) Test_emptyQueryString() {
	response := s.getMostPossibleMatches("")
	s.shouldResponseEmptyContent(response)
}

func (s *GetSinglesTestSuite) Test_nonPositiveMostPossibleQuery() {
	response := s.getMostPossibleMatches("0")
	s.shouldResponseEmptyMatches(response)
}

func (s *GetSinglesTestSuite) Test_noSingleExists() {
	response := s.getMostPossibleMatches("1")
	s.shouldResponseEmptyMatches(response)
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseAllShortestGirls() {
	s.givenSingleAdded("GIRL", 170, 1)
	s.givenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "GIRL")).
		Assert(jsonpath.Equal("$[0].height", float64(165))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responsePartialShortestGirls() {
	s.givenSingleAdded("GIRL", 165, 1)
	s.givenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "GIRL")).
		Assert(jsonpath.Equal("$[0].height", float64(165))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseAllHighestBoys() {
	s.givenSingleAdded("BOY", 170, 1)
	s.givenSingleAdded("BOY", 185, 1)

	response := s.getMostPossibleMatches("1")
	response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(185))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responsePartialHighestBoys() {
	s.givenSingleAdded("BOY", 170, 1)
	s.givenSingleAdded("BOY", 170, 1)

	response := s.getMostPossibleMatches("1")
	response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(170))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *GetSinglesTestSuite) givenSingleAdded(gender string, height int, wantedDates int) {
	single := &matching.Single{
		Gender:      gender,
		Height:      height,
		WantedDates: wantedDates,
	}

	request, err := json.Marshal(single)
	if err != nil {
		panic(err)
	}

	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(string(request)).
		Expect(s.T()).
		End()
}

func (s *GetSinglesTestSuite) getMostPossibleMatches(mostPossibleQuery string) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(s.Url).
		Query(router.QueryKeyMostPossible, mostPossibleQuery).
		Expect(s.T())
}

func (s *GetSinglesTestSuite) shouldResponseEmptyContent(response *apitest.Response) apitest.Result {
	return response.
		Status(http.StatusBadRequest).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func (s *GetSinglesTestSuite) shouldResponseEmptyMatches(response *apitest.Response) apitest.Result {
	return response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 0)).
		End()
}
