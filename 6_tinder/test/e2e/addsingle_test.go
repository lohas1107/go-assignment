package e2e

import (
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
	ShouldResponseBadRequest(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveHeight() {
	response := GivenSingleAdding("GIRL", -200, 0)
	ShouldResponseBadRequest(response)
}

func (s *AddSingleTestSuite) Test_nonPositiveWantedDates() {
	response := GivenSingleAdding("BOY", 200, 0)
	ShouldResponseBadRequest(response)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addBoy() {
	response := GivenSingleAdding("BOY", 180, 1)
	ShouldResponseEmptyMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_noSingleExists_addGirl() {
	response := GivenSingleAdding("GIRL", 170, 1)
	ShouldResponseEmptyMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addBoyButNoAnyMatch() {
	GivenSingleAdded("GIRL", 170, 1)
	response := GivenSingleAdding("BOY", 160, 1)
	ShouldResponseEmptyMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addGirlButNoAnyMatch() {
	GivenSingleAdded("BOY", 160, 1)
	response := GivenSingleAdding("GIRL", 170, 1)
	ShouldResponseEmptyMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addBoyAndMatch() {
	GivenSingleAdded("GIRL", 160, 1)
	response := GivenSingleAdding("BOY", 185, 1)
	assert := AssertMatchesLength(response, http.StatusCreated, 1)
	assert = AssertMatchesContent(assert, 0, "GIRL", 160, 1)
	assert.End()
}

func (s *AddSingleTestSuite) Test_addGirlAndMatch() {
	GivenSingleAdded("BOY", 185, 1)
	response := GivenSingleAdding("GIRL", 160, 1)
	assert := AssertMatchesLength(response, http.StatusCreated, 1)
	assert = AssertMatchesContent(assert, 0, "BOY", 185, 1)
	assert.End()
}
