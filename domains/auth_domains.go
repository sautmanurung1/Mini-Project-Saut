package domains

import (
	"Tugas-Mini-Project/entities"
)

type AuthRepository interface {
	Login(username, password string) (credential entities.User, err error)
	Register(user entities.User) error
}

type AuthService interface {
	LoginService(username, password string) (string, int)
	RegisterService(credential entities.User) error
}
