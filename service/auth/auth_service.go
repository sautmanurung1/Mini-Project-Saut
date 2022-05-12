package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/infrastructure/http/middleware"
	"net/http"
)

type svcAuth struct {
	c    database.Config
	repo domains.AuthRepository
}

func NewAuthService(repo domains.AuthRepository, c database.Config) domains.AuthService {
	return &svcAuth{
		repo: repo,
		c:    c,
	}
}

func (s *svcAuth) LoginService(username, password string, roleId int) (string, int) {
	user, _ := s.repo.Login(username, password, roleId)

	if (user.Password != password) || (user == entities.User{}) {
		return "Your Password Error", http.StatusUnauthorized
	}

	if user.RoleId == 1 {
		token, _ := middleware.CreateToken(int(user.ID), user.Username, s.c.Login_Teacher)
		return token, http.StatusOK
	} else {
		token, _ := middleware.CreateToken(int(user.ID), user.Username, s.c.Login_Student)
		return token, http.StatusOK
	}
}

func (s *svcAuth) RegisterService(credential entities.User) error {
	return s.repo.Register(credential)
}
