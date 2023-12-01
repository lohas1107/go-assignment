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
}

func (s *AddSingleTestSuite) SetupTest() {
	s.Url = GetUrl("/singles")

}

func (s *AddSingleTestSuite) Test_givenNoAnySingle_addOneBoy() {
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(`{
			"name": "Bob",
			"gender": "BOY",
			"height": 185,
			"wantedDates": 1
		}`).
		Expect(s.T()).
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}
