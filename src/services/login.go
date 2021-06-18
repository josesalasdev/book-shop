package services

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginModel struct {
	email    string
	password string
}

func (this *loginModel) LoginUser(email string, password string) bool {
	return this.email == email && this.password == password
}
