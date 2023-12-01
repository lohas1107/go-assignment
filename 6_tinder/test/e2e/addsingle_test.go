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
	body := `{ 
	"name": "Bob",
	"gender": "BOY",
	"height": 185,
	"wantedDates": 1
}`
	response := s.addSingle(body)
	response.
		Status(http.StatusCreated).
		Assert(jsonpath.Len("$", 0)).
		End()
}

func (s *AddSingleTestSuite) addSingle(body string) *apitest.Response {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Post(s.Url).
		Body(body).
		Expect(s.T())
}
