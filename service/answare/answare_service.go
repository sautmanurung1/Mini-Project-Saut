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

func (s *svcAnsware) GetAnswareByIdService(id int) (entities.Answare, error) {
	var answare entities.Answare
	if id != answare.QuestionId {
		return answare, nil
	} else {
		return s.repo.GetAnswareById(id)
	}
}

func (s *svcAnsware) GetAllAnswareService() (ans []entities.Answare, err error) {
	return s.repo.GetAllAnsware()
}

func (s *svcAnsware) UpdateAnswareService(id int, answare entities.Answare) (entities.Answare, error) {
	if answare.UserId == 1 && answare.QuestionId == id {
		return answare, nil
	} else {
		return s.repo.UpdateAnsware(id, answare)
	}
}

func (s *svcAnsware) DeleteAnswareService(id int, ans entities.Answare) (string, error) {
	if ans.UserId == 1 && ans.QuestionId == id {
		return "user can't delete the answare because user is teacher", nil
	} else {
		return "delete answare from question", s.repo.DeleteAnsware(id, ans)
	}
}
