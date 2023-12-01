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

func (s *AddSingleTestSuite) Test_givenNoAnySingle_addOneBoy() {
	body := s.givenSingle(&matching.Single{})
	response := s.addSingle(body)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}

func (s *AddSingleTestSuite) Test_() {
	body := s.givenSingle(&matching.Single{})
	response := s.addSingle(body)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}

func (s *AddSingleTestSuite) givenSingle(single *matching.Single) string {
	body, err := json.Marshal(single)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func (s *AddSingleTestSuite) addSingle(body string) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(body).
		Expect(s.T())
}
