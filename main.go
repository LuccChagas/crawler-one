package main

import (
	"log"
	"os"

	"github.com/LuccChagas/crawler-one/service"
)

func main() {
	log.Println("Welcome to Crawler App!")

	//TODO: need to recieve config
	app := service.NewCli()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
