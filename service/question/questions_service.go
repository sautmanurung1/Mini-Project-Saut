package question

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
)

type svcQuestion struct {
	c    database.Config
	repo domains.QuestionRepository
}

func NewQuestionService(repo domains.QuestionRepository, c database.Config) domains.QuestionService {
	return &svcQuestion{
		c:    c,
		repo: repo,
	}
}

func (s *svcQuestion) CreateQuestionService(question entities.Question) (string, error) {
	if question.UserId == 1 {
		return "teacher create question from assignment", s.repo.CreateQuestion(question)
	} else {
		return "user can't make question because the user is student", nil
	}
}

func (s *svcQuestion) GetQuestionByIDService(id int) (entities.Question, error) {
	return s.repo.GetQuestionByID(id)
}

func (s *svcQuestion) GetAllQuestionService() (question []entities.Question, err error) {
	return s.repo.GetAllQuestion()
}

func (s *svcQuestion) UpdateQuestionService(id int, question entities.Question) (entities.Question, error) {
	if question.UserId == 1 && question.AssignmentId == id {
		return s.repo.UpdateQuestion(id, question)
	} else {
		return question, nil
	}
}

func (s *svcQuestion) DeleteQuestionService(id int, question entities.Question) (string, error) {
	if question.UserId == 1 && question.AssignmentId == id {
		return "delete question by id", s.repo.DeleteQuestion(id, question)
	} else {
		return "user can't update question because user is student", nil
	}
}
