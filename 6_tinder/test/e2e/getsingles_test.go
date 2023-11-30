package e2e

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

const (
	Host    = "http://localhost:8080"
	Version = "v1"
)

func TestGetSinglesTestSuite(t *testing.T) {
	suite.Run(t, new(GetSinglesTestSuite))
}

type GetSinglesTestSuite struct {
	suite.Suite
	Url string
}

func (s *GetSinglesTestSuite) SetupTest() {
	s.Url = fmt.Sprintf("%s/%s%s", Host, Version, "/singles")
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
