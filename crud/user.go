package crud

import (
	"imdb-movie/ent"
	"log"
)

type UserStruct struct {
	Email    string
	Password string
	FullName string
	Username string
}

func (crud *Crud) GetAllUsers() ([]*ent.User, error) {
	if users, err := crud.Client.User.Query().All(crud.Ctx); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (crud *Crud) GetUser(id int) (*ent.User, error) {
	if u, err := crud.Client.User.Get(crud.Ctx, id); err != nil {
		log.Println("on GetUser Function in crud/user.go: ", err)
		return nil, err
	} else {
		log.Println("User Fetched Completly: ", u)
		return u, nil
	}
}

func (crud *Crud) AddUser(u *UserStruct) (*ent.User, error) {
	newUser, err := crud.Client.User.Create().SetEmail(u.Email).SetFullName(u.FullName).SetPassword(u.Password).SetUsername(u.Username).Save(crud.Ctx)
	if err != nil {
		log.Println("on AddUser Function in crud/user.go: ", err)
		return nil, err
	} else {
		log.Println("User Created: ", u)
		return newUser, err
	}
}

func (crud *Crud) AddMoviesToUser(ids *[]int, userID int) ([]*ent.Movie, error) {
	user, err := crud.Client.User.UpdateOneID(userID).AddMovieIDs(*ids...).Save(crud.Ctx)
	if err != nil {
		log.Println("on AddMoviesToUser Function in crud/user.go: ", err)
		return nil, err
	}
	movies, err := user.QueryMovies().All(crud.Ctx)
	if err != nil {
		log.Println("on AddMoviesToUser Function in crud/user.go: ", err)
		return nil, err
	}
	return movies, nil
}

func (crud *Crud) GetUserMovie(userID int) ([]*ent.Movie, error) {
	user, err := crud.Client.User.Get(crud.Ctx, userID)
	if err != nil {
		log.Println("on GetUserMovie Function in crud/user.go: ", err)
		return nil, err
	}
	movies, err := user.QueryMovies().All(crud.Ctx)

	if err != nil {
		log.Println("on GetUserMovie Function in crud/user.go: ", err)
		return nil, err
	}

	return movies, nil
}

func (crud *Crud) AddFavorites(ids *[]int, userID int) ([]*ent.Movie, error) {
	_, err := crud.Client.User.UpdateOneID(userID).AddFavoriteIDs(*ids...).Save(crud.Ctx)
	if err != nil {
		log.Println("on AddFavorites Function in crud/user.go: ", err)
		return nil, err
	}
	favorites, err := crud.GetFavorites(userID)
	if err != nil {
		log.Println("on AddFavorites Function in crud/user.go: ", err)
		return nil, err
	}
	return favorites, err
}

func (crud *Crud) GetFavorites(userID int) ([]*ent.Movie, error) {
	user, err := crud.GetUser(userID)
	if err != nil {
		log.Println("on GetFavorites Function in crud/user.go: ", err)
		return nil, err
	}
	favorites, err := user.QueryFavorites().All(crud.Ctx)
	if err != nil {
		log.Println("on GetFavorites Function in crud/user.go: ", err)
		return nil, err
	}
	return favorites, nil
}
