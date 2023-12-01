package e2e

import "fmt"

const (
	Host    = "http://localhost:8080"
	Version = "v1"
)

func GetUrl(path string) string {
	return fmt.Sprintf("%s/%s%s", Host, Version, path)
}
