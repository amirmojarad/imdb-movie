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

func (userRouter *API) GetUsers() []*ent.User {
	users, err := userRouter.Client.User.Query().Order(ent.Asc(user.FieldID)).All(userRouter.Ctx)
	if err != nil {
		log.Println("on GetUsers error occured: ", err)
	}
	return users
}
