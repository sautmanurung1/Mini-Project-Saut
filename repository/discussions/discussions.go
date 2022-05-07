package discussions

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewDiscussionsRepository(db *gorm.DB) domains.DiscussionsRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetAllDiscussions(discussions entities.Discussions) ([]entities.Discussions, error) {
	discuss := []entities.Discussions{}
	var user entities.User
	var questions entities.Question
	var answare entities.Answare

	r.DB.Find(&discuss)
	discussions.QuestionId = int(questions.ID)
	discussions.UserId = int(user.ID)
	discussions.AnswareId = int(answare.ID)
	discussions.Name = user.Name
	discussions.QuestionUser = questions.QuestionUser
	discussions.AnswereUser = answare.AnswareUser

	return discuss, nil
}
