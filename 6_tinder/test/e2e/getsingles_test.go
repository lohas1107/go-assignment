package e2e

import (
	"fmt"
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
	GivenSingleAdded("GIRL", 170, 1)
	GivenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "GIRL", 165, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responsePartialShortestGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 165, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "GIRL", 165, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseMultiShortGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 170, 1)

	response := s.getMostPossibleMatches("2")
	assert := s.assertPossibleMatchSize(response, 2)
	assert = s.assertResponseContent(assert, "0", "GIRL", 165, 1)
	assert = s.assertResponseContent(assert, "1", "GIRL", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noBoyExists_responseInsufficientShortGirls() {
	GivenSingleAdded("GIRL", 165, 1)
	GivenSingleAdded("GIRL", 170, 1)

	response := s.getMostPossibleMatches("3")
	assert := s.assertPossibleMatchSize(response, 2)
	assert = s.assertResponseContent(assert, "0", "GIRL", 165, 1)
	assert = s.assertResponseContent(assert, "1", "GIRL", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseAllHighestBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 185, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "BOY", 185, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responsePartialHighestBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseMultiHighBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 180, 1)

	response := s.getMostPossibleMatches("2")
	assert := s.assertPossibleMatchSize(response, 2)
	assert = s.assertResponseContent(assert, "0", "BOY", 180, 1)
	assert = s.assertResponseContent(assert, "1", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_noGirlExists_responseInsufficientHighBoys() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 180, 1)

	response := s.getMostPossibleMatches("3")
	assert := s.assertPossibleMatchSize(response, 2)
	assert = s.assertResponseContent(assert, "0", "BOY", 180, 1)
	assert = s.assertResponseContent(assert, "1", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseAllMostPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responsePartialPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("1")
	assert := s.assertPossibleMatchSize(response, 1)
	assert = s.assertResponseContent(assert, "0", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseMultiPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("2")
	assert := s.assertPossibleMatchSize(response, 2)
	assert = s.assertResponseContent(assert, "0", "GIRL", 180, 1)
	assert = s.assertResponseContent(assert, "1", "BOY", 170, 1)
	assert.End()
}

func (s *GetSinglesTestSuite) Test_BoysAndGirlsExist_responseInsufficientPossibleSingles() {
	GivenSingleAdded("BOY", 170, 1)
	GivenSingleAdded("GIRL", 175, 1)
	GivenSingleAdded("GIRL", 180, 1)
	GivenSingleAdded("GIRL", 180, 1)

	response := s.getMostPossibleMatches("5")
	assert := s.assertPossibleMatchSize(response, 4)
	assert = s.assertResponseContent(assert, "0", "BOY", 170, 1)
	assert = s.assertResponseContent(assert, "1", "GIRL", 175, 1)
	assert = s.assertResponseContent(assert, "2", "GIRL", 180, 1)
	assert = s.assertResponseContent(assert, "3", "GIRL", 180, 1)
	assert.End()
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
	return s.assertPossibleMatchSize(response, 0).
		End()
}

func (s *GetSinglesTestSuite) assertPossibleMatchSize(response *apitest.Response, length int) *apitest.Response {
	return response.
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", length))
}

func (s *GetSinglesTestSuite) assertResponseContent(
	assert *apitest.Response,
	index string,
	gender string,
	height int,
	wantedDates int,
) *apitest.Response {
	return assert.
		Assert(jsonpath.Equal(fmt.Sprintf("$[%s].gender", index), gender)).
		Assert(jsonpath.Equal(fmt.Sprintf("$[%s].height", index), float64(height))).
		Assert(jsonpath.Equal(fmt.Sprintf("$[%s].wantedDates", index), float64(wantedDates)))
}
