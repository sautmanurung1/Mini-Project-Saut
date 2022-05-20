package domains

import (
	"Tugas-Mini-Project/entities"
)

type AuthRepository interface {
	Login(username string) (user entities.User, err error)
	Register(user entities.User) error
}

type AuthService interface {
	LoginService(username string, password string, roleId int) (string, int)
	RegisterService(credential entities.User) error
}
