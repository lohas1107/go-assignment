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
}

func (s *AddSingleTestSuite) SetupTest() {
	s.Url = GetUrl("/singles")
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Delete(s.Url).
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}

func (s *AddSingleTestSuite) Test_invalidGender() {
	invalidSingle := &matching.Single{
		Gender: "",
		Height: 0,
	}

	response := s.givenAddedSingle(invalidSingle)
	s.shouldNotResponseContent(response, http.StatusBadRequest)
}

func (s *AddSingleTestSuite) Test_nonPositiveHeight() {
	invalidSingle := &matching.Single{
		Gender: "GIRL",
		Height: -200,
	}

	response := s.givenAddedSingle(invalidSingle)
	s.shouldNotResponseContent(response, http.StatusBadRequest)
}

func (s *AddSingleTestSuite) Test_givenNoAnySingle_addOneBoy() {
	boy := &matching.Single{
		Gender: "BOY",
		Height: 180,
	}

	response := s.givenAddedSingle(boy)
	s.shouldNotResponseMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addBoyButNoAnyMatch() {
	girl := &matching.Single{
		Gender: "GIRL",
		Height: 170,
	}

	response := s.givenAddedSingle(girl)
	s.shouldNotResponseMatches(response, http.StatusCreated)

	boy := &matching.Single{
		Gender:      "BOY",
		Height:      160,
		WantedDates: 1,
	}

	response = s.givenAddedSingle(boy)
	s.shouldNotResponseMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addGirlButNoAnyMatch() {
	boy := &matching.Single{
		Gender:      "BOY",
		Height:      160,
		WantedDates: 1,
	}

	response := s.givenAddedSingle(boy)
	s.shouldNotResponseMatches(response, http.StatusCreated)

	girl := &matching.Single{
		Gender: "GIRL",
		Height: 170,
	}

	response = s.givenAddedSingle(girl)
	s.shouldNotResponseMatches(response, http.StatusCreated)
}

func (s *AddSingleTestSuite) Test_addAndMatch() {
	boy := &matching.Single{
		Gender:      "BOY",
		Height:      185,
		WantedDates: 1,
	}

	response := s.givenAddedSingle(boy)
	s.shouldNotResponseMatches(response, http.StatusCreated)

	girl := &matching.Single{
		Gender: "GIRL",
		Height: 160,
	}

	response = s.givenAddedSingle(girl)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(185))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) givenAddedSingle(single *matching.Single) *apitest.Response {
	request, err := json.Marshal(single)
	if err != nil {
		panic(err)
	}

	response := apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(string(request)).
		Expect(s.T())

	return response
}

func (s *AddSingleTestSuite) shouldNotResponseContent(response *apitest.Response, created int) apitest.Result {
	return response.
		Status(created).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func (s *AddSingleTestSuite) shouldNotResponseMatches(response *apitest.Response, created int) apitest.Result {
	return response.
		Status(created).
		Assert(jsonpath.Len("$", 0)).
		End()
}
