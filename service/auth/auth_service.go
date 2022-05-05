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

func (s *svcAuth) LoginService(credential entities.User) (string, int) {
	_ = s.repo.Login(credential)

	if credential.Username != credential.Password {
		return "Username atau Password salah", http.StatusUnauthorized
	}

	token, err := middleware.CreateToken(int(credential.ID), credential.Username, s.c.Token)
	if err != nil {
		return "Token tidak dapat dibuat", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svcAuth) RegisterService(credential entities.User) error {
	/*
		data := entities.User{}
		user := s.repo.Register(data)
		return user
	*/
	var role entities.Role
	if credential.RoleId == 1 {
		e := entities.User{
			Username: credential.Username,
			Password: credential.Password,
			RoleId:   role.ID,
			Role:     role,
		}
		err := s.repo.Register(e)
		if err != nil {
			return err
		}
	} else if credential.RoleId == 2 {
		e := entities.User{
			Username: credential.Username,
			Password: credential.Password,
			RoleId:   role.ID,
			Role:     role,
		}
		err := s.repo.Register(e)
		if err != nil {
			return err
		}
	}
	return nil
}
