package answare

import (
	"Tugas-Mini-Project/domains/answare"
	"Tugas-Mini-Project/domains/question"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAnswareRepository(db *gorm.DB) answare.AnswareRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateAnsware(answare answare.Answare) error {
	var questions question.Question
	r.DB.Where("id = ?", answare.QuestionId).First(&questions)
	answare.Questions = questions.QuestionUser
	answare.Name = questions.Name
	r.DB.Create(&answare)
	return nil
}

func (r *repository) GetAnswareById(id int) (answare.Answare, error) {
	var questions question.Question
	var ans answare.Answare
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&ans)
	r.DB.Where("id = ?", ans.QuestionId).First(&questions)
	ans.Name = questions.Name
	ans.Questions = questions.QuestionUser
	return ans, nil
}

func (r *repository) GetAllAnsware() ([]answare.Answare, error) {
	ans := []answare.Answare{}
	r.DB.Find(&ans)
	return ans, nil
}

func (r *repository) UpdateAnsware(id int, answare answare.Answare) (answare.Answare, error) {
	var questions question.Question
	r.DB.Model(&answare).Where("id = ?", id).Updates(&answare)
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&answare)
	r.DB.Where("id = ?", answare.QuestionId).First(&questions)
	answare.Name = questions.Name
	answare.Questions = questions.QuestionUser
	return answare, nil
}

func (r *repository) DeleteAnsware(id int) (answare.Answare, error) {
	var ans answare.Answare
	r.DB.Where("id = ?", id).Delete(&ans)
	return ans, nil
}
