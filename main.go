package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vincen320/user-service-graphql/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
