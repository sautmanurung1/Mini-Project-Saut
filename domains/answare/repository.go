package answare

import "Tugas-Mini-Project/entities"

type AnswareRepository interface {
	CreateAnsware(answare entities.Answare) error
	GetAnswareById(id int) (entities.Answare, error)
	GetAllAnsware() ([]entities.Answare, error)
	UpdateAnsware(id int, answare entities.Answare) (entities.Answare, error)
	DeleteAnsware(id int) (entities.Answare, error)
}
