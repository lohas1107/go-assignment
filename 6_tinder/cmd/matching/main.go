package main

import (
	"tinder/cmd/matching/router"
)

func main() {
	server := router.SetUp()
	server.Run()
}
