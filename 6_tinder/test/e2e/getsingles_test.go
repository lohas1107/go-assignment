package e2e

import (
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"tinder/cmd/matching/router"
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
	Reset(s.T())
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
