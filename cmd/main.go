package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	apiservice "github.com/pradeep/aws/cmd/api"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	server := apiservice.NewApiServer(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
