package api

import (
	"imdb-movie/ent"
	"imdb-movie/ent/user"
	"log"
)

func (userRouter *API) GetUser(id int) (*ent.User, error) {
	fetchedUser, err := userRouter.Client.User.Get(userRouter.Ctx, id)
	if err != nil {
		log.Println("error when fetching user by id: ", err)
		return nil, err
	} else if fetchedUser == nil {
		log.Println("no user with id: ", id)
		return nil, nil
	}
	return fetchedUser, nil
}

func (userRouter *API) PostUsers() (*ent.User, error) {
	// newUser, err := userRouter.
	// 			Client.User.Create().
	// 			SetUsername().
	// 			Save(userRouter.Ctx)

	return nil, nil
}

func (userRouter *API) GetUserMovies(id int) ([]*ent.Movie, error) {
	user, err := userRouter.Client.User.Get(userRouter.Ctx, id)
	log.Println(user)
	if err != nil {
		return nil, err
	}
	movies, err := user.QueryMovies().All(userRouter.Ctx)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (userRouter *API) AddMovies2User(user *ent.User, moviesIDs []int) {
	for _, id := range moviesIDs {
		userRouter.Client.User.UpdateOne(user).AddMovieIDs(id).Save(userRouter.Ctx)
	}
}

func (userRouter *API) GetUsers() []*ent.User {
	users, err := userRouter.Client.User.Query().Order(ent.Asc(user.FieldID)).All(userRouter.Ctx)
	if err != nil {
		log.Println("on GetUsers error occured: ", err)
	}
	return users
}
