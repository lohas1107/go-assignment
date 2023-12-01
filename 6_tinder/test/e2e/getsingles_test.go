package e2e

import (
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
}

func (s *GetSinglesTestSuite) SetupTest() {
	s.Url = GetUrl("/singles")
	Reset(s.T())
}

func (s *GetSinglesTestSuite) Test_emptyQueryString() {
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(s.Url).
		Expect(s.T()).
		Status(http.StatusBadRequest).
		Assert(jsonpath.NotPresent("$")).
		End()
}

func (s *GetSinglesTestSuite) Test_nonPositiveMostPossibleQuery() {
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(s.Url).
		Query(router.QueryKeyMostPossible, "0").
		Expect(s.T()).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$", 0)).
		End()
}
