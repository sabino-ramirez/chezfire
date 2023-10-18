package main

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sabino-ramirez/chezfire/auth"
	"github.com/sabino-ramirez/chezfire/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("env loading err:", err)
	}

	// log.Println(os.Getenv("SPOTIFY_ID"))
	if err := run(); err != nil {
		// fmt.Fprintf(os.Stderr, "%s\n", err)
		log.Println("run() err:", err)
	}

}

func run() error {
	ctx := context.Background()
	clientCredsHttpClient := auth.NewClientCredHelper().Config.Client(ctx)

	srv := server.NewServer(clientCredsHttpClient)

	portNumber, portNumberExists := os.LookupEnv("PORT")
	if !portNumberExists {
		portNumber = "8000"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", portNumber), srv); err != nil {
		log.Fatal(err)
	}

	return nil
}
