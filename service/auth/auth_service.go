package auth

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
	"Tugas-Mini-Project/infrastructure/http/middleware"
	"golang.org/x/crypto/bcrypt"
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

func (s *svcAuth) LoginService(username string, password string, roleId int) (string, int) {
	user, _ := s.repo.Login(username)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if er != nil {
		return "wrong password", http.StatusUnauthorized
	}

	if user.RoleId != roleId {
		return "your role id error", http.StatusUnauthorized
	}

	token, _ := middleware.CreateToken(int(user.ID), user.Username, s.c.JWT_KEY)
	return token, http.StatusOK
}

func (s *svcAuth) RegisterService(credential entities.User) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(credential.Password), bcrypt.MinCost)
	credential.Password = string(password)
	return s.repo.Register(credential)
}
