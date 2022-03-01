package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ResponseModel struct {
	QueryString string
	Results     []MovieData
}

type MovieData struct {
	Image       string
	Title       string
	Description string
	RunTimeStr  string
	Genres      string
	ImdbRating  string
	Plot        string
	Stars       string
}

type IMDBQueryData struct {
	Title string
	Sort  string
	Count int
}

func generateQuery(title, sort string, count int) string {
	apiKey := os.Getenv("API_KEY")
	return fmt.Sprintf("https://imdb-api.com/API/AdvancedSearch/%s?title=%s&count=%d&sort=%s", apiKey, title, count, sort)

}

func GETMovies(title string) (*[]MovieData, error) {
	resp, err := http.Get(generateQuery(title, "", 100))
	if err != nil {
		log.Println("on GETMovies in api/imdb_router.go: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var movieData ResponseModel
	err = decoder.Decode(&movieData)
	if err != nil {
		log.Println("on GETMovies in api/imdb_router.go: ", err)
		return nil, err
	}
	return &movieData.Results, nil
}
