package api

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

func (movieRouter *API) AddMovies(movies *[]MovieData) {
	for _, m := range *movies {
		movieRouter.Client.Movie.Create().
			SetGenre(m.Genres).
			SetPoster(m.Image).
			SetPlot(m.Plot).
			SetStars(m.Plot).
			SetImdbRating(m.ImdbRating).
			SetDescription(m.Description).Save(movieRouter.Ctx)
	}
}

func GetMovie(movieTitle string) []MovieData {
	api_key := os.Getenv("API_KEY")
	log.Println(api_key)
	resp, err := http.Get(fmt.Sprintf("https://imdb-api.com/API/AdvancedSearch/%s?title=%s", api_key, movieTitle))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var data ResponseModel
	_ = decoder.Decode(&data)

	return data.Results
}
