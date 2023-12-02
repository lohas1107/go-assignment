package e2e

import (
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"tinder/cmd/matching/router"
)

type GetSinglesTestSuite struct {
	suite.Suite
}

func TestGetSinglesTestSuite(t *testing.T) {
	suite.Run(t, new(GetSinglesTestSuite))
	test = t
}

func (s *GetSinglesTestSuite) SetupTest() {
	Reset()
}

func (s *GetSinglesTestSuite) Test_emptyQueryString() {
	response := s.getMostPossibleMatches("")
	ShouldResponseBadRequest(response)
}

func (s *GetSinglesTestSuite) Test_nonPositiveMostPossibleQuery() {
	response := s.getMostPossibleMatches("0")
	ShouldResponseEmptyMatches(response, http.StatusOK)
}

func (s *GetSinglesTestSuite) Test_noSingleExists() {
	response := s.getMostPossibleMatches("1")
	ShouldResponseEmptyMatches(response, http.StatusOK)
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseAllShortestGirls() {
	GivenSingleAdded("GIRL", 170, 1)
	GivenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "GIRL", 165, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responsePartialShortestGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "GIRL", 165, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseMultiShortGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 170, 1)

	response := s.getMostPossibleMatches("2")
	assert := AssertMatchesLength(response, http.StatusOK, 2)
	assert = AssertMatchesContent(assert, 0, "GIRL", 165, 1)
	assert = AssertMatchesContent(assert, 1, "GIRL", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseInsufficientShortGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 170, 1)

	response := s.getMostPossibleMatches("3")
	assert := AssertMatchesLength(response, http.StatusOK, 2)
	assert = AssertMatchesContent(assert, 0, "GIRL", 165, 1)
	assert = AssertMatchesContent(assert, 1, "GIRL", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseAllHighestBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 185, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "BOY", 185, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responsePartialHighestBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseMultiHighBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 180, 1)

	response := s.getMostPossibleMatches("2")
	assert := AssertMatchesLength(response, http.StatusOK, 2)
	assert = AssertMatchesContent(assert, 0, "BOY", 180, 1)
	assert = AssertMatchesContent(assert, 1, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseInsufficientHighBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 180, 1)

	response := s.getMostPossibleMatches("3")
	assert := AssertMatchesLength(response, http.StatusOK, 2)
	assert = AssertMatchesContent(assert, 0, "BOY", 180, 1)
	assert = AssertMatchesContent(assert, 1, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseAllMostPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responsePartialPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("1")
	assert := AssertMatchesLength(response, http.StatusOK, 1)
	assert = AssertMatchesContent(assert, 0, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseMultiPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("2")
	assert := AssertMatchesLength(response, http.StatusOK, 2)
	assert = AssertMatchesContent(assert, 0, "GIRL", 180, 1)
	assert = AssertMatchesContent(assert, 1, "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseInsufficientPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 175, 1)
	GivenSingleAdded("GIRL", 180, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("5")
	assert := AssertMatchesLength(response, http.StatusOK, 4)
	assert = AssertMatchesContent(assert, 0, "BOY", 170, 1)
	assert = AssertMatchesContent(assert, 1, "GIRL", 175, 1)
	assert = AssertMatchesContent(assert, 2, "GIRL", 180, 1)
	assert = AssertMatchesContent(assert, 3, "GIRL", 180, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) getMostPossibleMatches(mostPossibleQuery string) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(s.Url).
		Query(router.QueryKeyMostPossible, mostPossibleQuery).
		Expect(s.T())
}
