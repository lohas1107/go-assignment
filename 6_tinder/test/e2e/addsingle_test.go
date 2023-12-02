package e2e

import (
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type AddSingleTestSuite struct {
	suite.Suite
	httpClient *http.Client
	Url        string
}

func TestAddSingleTestSuite(t *testing.T) {
	suite.Run(t, new(AddSingleTestSuite))
	test = t
}

func (s *AddSingleTestSuite) SetupTest() {
	s.Url = GetUrl("/singles")
	Reset()
}

func (s *AddSingleTestSuite) Test_invalidGender() {
	response := GivenSingleAdding("", 0, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveHeight() {
	response := GivenSingleAdding("GIRL", -200, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveWantedDates() {
	response := GivenSingleAdding("BOY", 200, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addBoy() {
	response := GivenSingleAdding("BOY", 180, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addGirl() {
	response := GivenSingleAdding("GIRL", 170, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addBoyButNoAnyMatch() {
	GivenSingleAdded("GIRL", 170, 1)

	response := GivenSingleAdding("BOY", 160, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addGirlButNoAnyMatch() {
	GivenSingleAdded("BOY", 160, 1)

	response := GivenSingleAdding("GIRL", 170, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addBoyAndMatch() {
	GivenSingleAdded("GIRL", 160, 1)

	response := GivenSingleAdding("BOY", 185, 1)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "GIRL")).
		Assert(jsonpath.Equal("$[0].height", float64(160))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) Test_addGirlAndMatch() {
	GivenSingleAdded("BOY", 185, 1)

	response := GivenSingleAdding("GIRL", 160, 1)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(185))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) shouldResponseEmptyContent(response *apitest.Response) apitest.Result {
	return response.
		Status(http.StatusBadRequest).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func (s *AddSingleTestSuite) shouldResponseEmptyMatches(response *apitest.Response) apitest.Result {
	return response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}
