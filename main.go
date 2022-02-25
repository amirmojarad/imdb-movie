package main

import (
	"context"
	"fmt"
	"imdb-movie/api"
	"imdb-movie/ent"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

}

func main() {
	username, password := os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_USERNAME")
	client, err := ent.Open("postgres", fmt.Sprintf("user=%s dbname=imdb-movie password=%s sslmode=disable", username, password))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	client = client.Debug()

	log.Println("Database Connected!")

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	// Dump migration changes to stdout.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	router := api.API{Ctx: ctx, Client: client}
	router.RunAPI()
}
