package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
)

type svcAnsware struct {
	c    database.Config
	repo domains.AnswareRepository
}

func NewServiceAnsware(repo domains.AnswareRepository, c database.Config) domains.AnswareService {
	return &svcAnsware{
		repo: repo,
		c:    c,
	}
}

func (s *svcAnsware) CreateAnswareService(answare entities.Answare) (string, error) {
	if answare.UserId == 1 {
		return "User Can't Answare Because the user is Teacher", nil
	} else {
		return "Student Create Answare From Question ...", s.repo.CreateAnsware(answare)
	}
}

func (s *svcAnsware) GetAnswareByIdService(id int, answare entities.Answare) (string, error) {
	if id != answare.QuestionId {
		return "ID not Match", nil
	} else {
		return "Get Answare By Id ...", s.repo.GetAnswareById(id, answare)
	}
}

func (s *svcAnsware) GetAllAnswareService() (ans []entities.Answare, err error) {
	return s.repo.GetAllAnsware()
}

func (s *svcAnsware) UpdateAnswareService(id int, answare entities.Answare) (string, error) {
	if answare.UserId == 1 && answare.QuestionId != id {
		return "User Can't Answare Because the user is Teacher", nil
	} else {
		return "Student Update Answare From Question ...", s.repo.UpdateAnsware(id, answare)
	}
}

func (s *svcAnsware) DeleteAnswareService(id int, ans entities.Answare) (string, error) {
	return "Delete Answare From Question ...", s.repo.DeleteAnsware(id, ans)
}
