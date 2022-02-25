package main

import (
	"context"
	"imdb-movie/api"
	"imdb-movie/ent"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "user=postgres dbname=imdb-movie password=100990729 sslmode=disable")

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
	router := api.API{Ctx: ctx, Client: client, UserRouter: api.UserRouter{Ctx: ctx, Client: client}}
	router.RunAPI()
}
