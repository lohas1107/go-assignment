package e2e

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

const (
	Host    = "http://localhost:8080"
	Version = "v1"
)

func TestMatchingTestSuite(t *testing.T) {
	suite.Run(t, new(MatchingTestSuite))
}

type MatchingTestSuite struct {
	suite.Suite
}

func (s *MatchingTestSuite) SetupTest() {
}

func (s *MatchingTestSuite) Test_() {
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Get(fmt.Sprintf("%s/%s%s", Host, Version, "/singles")).
		Expect(s.T()).
		Status(http.StatusOK).
		End()
}
