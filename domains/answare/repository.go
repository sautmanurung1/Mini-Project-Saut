package answare

type AnswareRepository interface {
	CreateAnsware(answare Answare) error
	GetAnswareById(id int) (Answare, error)
	GetAllAnsware() ([]Answare, error)
	UpdateAnsware(id int, answare Answare) (Answare, error)
	DeleteAnsware(id int) (Answare, error)
}
