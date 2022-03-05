package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "test@testmail.com",
		password: "testpassword",
	}
}

func (info *loginInformation) LoginUser(email, password string) bool {
	return info.email == email && info.password == password
}
