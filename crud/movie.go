package crud

import (
	"imdb-movie/api/utils"
	"imdb-movie/ent"
	"log"
)

func (crud Crud) GetAllMovies() ([]*ent.Movie, error) {
	movies, err := crud.Client.Movie.Query().All(crud.Ctx)
	if err != nil {
		log.Println("on GetAllMovies Function in crud/movie.go: ", err)
		return nil, err
	}
	return movies, nil
}

func (crud Crud) GetMovieByID(id int) (*ent.Movie, error) {
	movie, err := crud.Client.Movie.Get(crud.Ctx, id)
	if err != nil {
		log.Println("on GetMovieByID Function in crud/movie.go: ", err)
		return nil, err
	}
	return movie, nil
}

func (crud Crud) AddAllMovies(movies *[]utils.MovieData) ([]*ent.Movie, error) {
	bulk := make([]*ent.MovieCreate, len(*movies))
	for i, m := range *movies {
		bulk[i] = crud.Client.Movie.Create().
			SetDescription(m.Description).
			SetGenre(m.Genres).
			SetImdbRating(m.ImdbRating).
			SetPlot(m.Plot).
			SetPoster(m.Image).
			SetStars(m.Stars).
			SetTitle(m.Title).SetRated(12.2)
	}
	if result, err := crud.Client.Movie.CreateBulk(bulk...).Save(crud.Ctx); err != nil {
		return nil, err
	} else {
		return result, err
	}
}
