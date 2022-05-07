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

func (r *repository) CreateDiscussions(discussions entities.Discussions) error {
	var user entities.User
	var questions entities.Question
	var answare entities.Answare
	r.DB.Where("id = ?", discussions.QuestionId).Find(&questions)
	r.DB.Where("id = ? ", discussions.UserId).Find(&user)
	r.DB.Where("id = ?", discussions.AnswareId).Find(&answare)
	discussions.Name = user.Name
	discussions.QuestionUser = questions.QuestionUser
	discussions.AnswereUser = answare.AnswareUser

	r.DB.Create(&discussions)
	return nil
}
func (r *repository) GetAllDiscussions() (discussions []entities.Discussions, err error) {
	r.DB.Find(&discussions)
	return discussions, nil
}
