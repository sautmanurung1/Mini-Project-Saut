package domains

import (
	"Tugas-Mini-Project/entities"
)

type AnswareRepository interface {
	CreateAnsware(answare entities.Answare) error
	GetAnswareById(id int) (entities.Answare, error)
	GetAllAnsware() (ans []entities.Answare, err error)
	UpdateAnsware(id int, answare entities.Answare) (entities.Answare, error)
	DeleteAnsware(id int, ans entities.Answare) error
}

type AnswareService interface {
	CreateAnswareService(answare entities.Answare) (string, error)
	GetAnswareByIdService(id int) (entities.Answare, error)
	GetAllAnswareService() (ans []entities.Answare, err error)
	UpdateAnswareService(id int, answare entities.Answare) (entities.Answare, error)
	DeleteAnswareService(id int, ans entities.Answare) (string, error)
}
