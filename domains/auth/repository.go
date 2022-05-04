package auth

import "Tugas-Mini-Project/entities"

type AuthRepository interface {
	Login(credential entities.User) error
	Register(user entities.User) error
}
