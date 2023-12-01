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
	body := s.givenSingle(&matching.Single{
		Gender: "",
		Height: 0,
	})
	response := s.addSingle(body)
	response.
		Status(http.StatusBadRequest).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func (s *AddSingleTestSuite) Test_givenNoAnySingle_addOneBoy() {
	body := s.givenSingle(&matching.Single{
		Gender: "BOY",
		Height: 0,
	})
	response := s.addSingle(body)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}

func (s *AddSingleTestSuite) Test_addAndMatch() {
	boy := &matching.Single{
		Gender:      "BOY",
		Height:      185,
		WantedDates: 1,
	}
	s.givenAddedSingle(boy)
	girl := &matching.Single{
		Gender: "GIRL",
		Height: 0,
	}
	response := s.givenAddedSingle(girl)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 1)).
		Assert(jsonpath.Equal("$[0].gender", "BOY")).
		Assert(jsonpath.Equal("$[0].height", float64(185))).
		Assert(jsonpath.Equal("$[0].wantedDates", float64(1))).
		End()
}

func (s *AddSingleTestSuite) givenSingle(single *matching.Single) string {
	body, err := json.Marshal(single)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func (s *AddSingleTestSuite) givenAddedSingle(single *matching.Single) *apitest.Response {
	body := s.givenSingle(single)
	response := s.addSingle(body)
	response.End()
	return response
}

func (s *AddSingleTestSuite) addSingle(body string) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(body).
		Expect(s.T())
}
