package main

import (
	"log"

	"github.com/LuccChagas/crawler-one/service"
)

func main() {
	log.Println("Welcome to Crawler App!")

	//TODO: need to recieve config
	err := service.NewCli()
	if err != nil {
		log.Fatal(err)
	}

}
