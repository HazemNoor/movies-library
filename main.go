package main

import (
	"github.com/HazemNoor/movies-library/app"
	"github.com/HazemNoor/movies-library/infrastructure"
	"log"
)

func init() {
	if err := infrastructure.LoadEnvVariables(); err != nil {
		log.Fatal(err)
	}

	if _, err := infrastructure.DbConnection(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := app.Dispatch(); err != nil {
		log.Fatal(err)
	}
}
