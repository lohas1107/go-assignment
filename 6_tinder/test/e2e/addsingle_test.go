package e2e

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"tinder/internal/matching"
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
	response := s.givenSingleAdded("", 0, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveHeight() {
	response := s.givenSingleAdded("GIRL", -200, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveWantedDates() {
	response := s.givenSingleAdded("BOY", 200, 0)
	s.shouldResponseEmptyContent(response)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addBoy() {
	response := s.givenSingleAdded("BOY", 180, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addGirl() {
	response := s.givenSingleAdded("GIRL", 170, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addBoyButNoAnyMatch() {
	response := s.givenSingleAdded("GIRL", 170, 1)
	s.shouldResponseEmptyMatches(response)

	response = s.givenSingleAdded("BOY", 160, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addGirlButNoAnyMatch() {
	response := s.givenSingleAdded("BOY", 160, 1)
	s.shouldResponseEmptyMatches(response)

	response = s.givenSingleAdded("GIRL", 170, 1)
	s.shouldResponseEmptyMatches(response)
}

func (s *AddSingleTestSuite) Test_addBoyAndMatch() {
	response := s.givenSingleAdded("GIRL", 160, 1)
	s.shouldResponseEmptyMatches(response)

	response = s.givenSingleAdded("BOY", 185, 1)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "GIRL")).
		Assert(jsonpath.Equal("$[0].height", float64(160))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) Test_addGirlAndMatch() {
	response := s.givenSingleAdded("BOY", 185, 1)
	s.shouldResponseEmptyMatches(response)

	response = s.givenSingleAdded("GIRL", 160, 1)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(185))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) givenSingleAdded(gender string, height int, wantedDates int) *apitest.Response {
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
		Post(s.Url).
		Body(string(request)).
		Expect(s.T())
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
