package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load the env file", err)
	}
}

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()
}
