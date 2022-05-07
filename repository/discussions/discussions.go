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

func (r *repository) GetAllDiscussions() ([]entities.Discussions, error) {
	var discussions []entities.Discussions
	err := r.DB.Find(&discussions).Error
	if err != nil {
		return nil, err
	}
	return discussions, nil
}
