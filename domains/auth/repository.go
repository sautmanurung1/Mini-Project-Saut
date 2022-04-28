package auth

import "Tugas-Mini-Project/model"

type AuthRepository interface {
	Login(credential model.User) (model.User, error)
	Register(user model.User) error
}
