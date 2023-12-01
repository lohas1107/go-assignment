package e2e

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

const (
	Host    = "http://localhost:8080"
	Version = "v1"
)

func GetUrl(path string) string {
	return fmt.Sprintf("%s/%s%s", Host, Version, path)
}

func Reset(t *testing.T) apitest.Result {
	return apitest.New().Debug().
		EnableNetworking(http.DefaultClient).
		Delete(GetUrl("/singles")).
		Expect(t).
		Status(http.StatusOK).
		End()
}
