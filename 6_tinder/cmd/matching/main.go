package main

import (
	"log"
	"tinder/cmd/matching/router"
)

func main() {
	server := router.SetUp()
	log.Fatalln(server.Run())
}
