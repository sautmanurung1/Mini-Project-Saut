package auth

type AuthRepository interface {
	Login(credential User) error
	Register(user User) error
}
