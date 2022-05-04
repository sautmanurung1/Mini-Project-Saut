package answare

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAnswareRepository(db *gorm.DB) domains.AnswareRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateAnsware(answare entities.Answare) error {
	var questions entities.Question
	r.DB.Where("id = ?", answare.QuestionId).Preload("Question").Find(&questions)
	answare.Questions = questions.QuestionUser
	answare.Name = questions.Name
	r.DB.Create(&answare)
	return nil
}

func (r *repository) GetAnswareById(id int) (entities.Answare, error) {
	var questions entities.Question
	var ans entities.Answare
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&ans)
	r.DB.Where("id = ?", ans.QuestionId).Preload("Question").Find(&questions)
	ans.Name = questions.Name
	ans.Questions = questions.QuestionUser
	return ans, nil
}

func (r *repository) GetAllAnsware() ([]entities.Answare, error) {
	ans := []entities.Answare{}
	r.DB.Find(&ans)
	return ans, nil
}

func (r *repository) UpdateAnsware(id int, answare entities.Answare) (entities.Answare, error) {
	var questions entities.Question
	r.DB.Model(&answare).Where("id = ?", id).Updates(&answare)
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&answare)
	r.DB.Where("id = ?", answare.QuestionId).First(&questions)
	answare.Name = questions.Name
	answare.Questions = questions.QuestionUser
	return answare, nil
}

func (r *repository) DeleteAnsware(id int) (entities.Answare, error) {
	var ans entities.Answare
	r.DB.Where("id = ?", id).Delete(&ans)
	return ans, nil
}
