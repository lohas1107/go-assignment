package e2e

import (
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type RemoveSingleTestSuite struct {
	suite.Suite
}

func TestRemoveSingleTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveSingleTestSuite))
	test = t
}

func (s *RemoveSingleTestSuite) SetupTest() {
	Reset()
}

func (s *RemoveSingleTestSuite) Test_removeSingle() {
	apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Delete(GetUrl("/singles")).
		Query("name", "someone").
		Expect(test).
		Status(http.StatusOK).
		End()
}
