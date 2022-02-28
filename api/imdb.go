package api

// import (
// 	"encoding/json"
// 	"fmt"
// 	"imdb-movie/ent"
// 	"imdb-movie/ent/movie"
// 	"log"
// 	"net/http"
// 	"os"
// )

// type ResponseModel struct {
// 	QueryString string
// 	Results     []MovieData
// }

// type MovieData struct {
// 	Image       string
// 	Title       string
// 	Description string
// 	RunTimeStr  string
// 	Genres      string
// 	ImdbRating  string
// 	Plot        string
// 	Stars       string
// }

// func (movieRouter *API) AddMovies(movies *[]MovieData) ([]*ent.Movie, error) {
// 	bulk := make([]*ent.MovieCreate, len(*movies))
// 	for i, m := range *movies {
// 		bulk[i] = movieRouter.Client.Movie.Create().
// 			SetDescription(m.Description).
// 			SetGenre(m.Genres).
// 			SetImdbRating(m.ImdbRating).
// 			SetPlot(m.Plot).
// 			SetPoster(m.Image).
// 			SetStars(m.Stars).
// 			SetTitle(m.Title).SetRated(12.2)
// 	}
// 	if result, err := movieRouter.Client.Movie.CreateBulk(bulk...).Save(movieRouter.Ctx); err != nil {
// 		return nil, err
// 	} else {
// 		return result, err
// 	}

// }

// func (api *API) GetMovie(id int) (*ent.Movie, error) {
// 	movie, err := api.Client.Movie.Get(api.Ctx, id)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return movie, nil
// }

// func (api *API) GetAllMovies() ([]*ent.Movie, error) {
// 	movies, err := api.Client.Movie.Query().Select(movie.Columns...).All(api.Ctx)
// 	log.Println(movies)
// 	if err != nil {
// 		log.Println("on GetAllMovies: ", err)
// 		return nil, err
// 	}
// 	return movies, nil
// }
// func GetMovie(movieTitle string) []MovieData {
// 	api_key := os.Getenv("API_KEY")
// 	log.Println(api_key)
// 	resp, err := http.Get(fmt.Sprintf("https://imdb-api.com/API/AdvancedSearch/%s?title=%s", api_key, movieTitle))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	decoder := json.NewDecoder(resp.Body)

// 	var data ResponseModel
// 	_ = decoder.Decode(&data)

// 	return data.Results
// }
