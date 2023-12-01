package main

import (
	"log"
	"tinder/cmd/matching/router"
	"tinder/internal/matching"
)

func main() {
	matching.Initialize()
	server := router.SetUp()
	log.Fatalln(server.Run())
}
