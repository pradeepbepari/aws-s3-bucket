package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pradeepbepari/aws-cloud/cmd/api"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	serve := api.NewApiServer(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err := serve.Run(); err != nil {
		log.Fatal(err)
	}
}
