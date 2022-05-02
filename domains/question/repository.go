package question

type QuestionRepository interface {
	CreateQuestion(question Question) error
	GetQuestionByID(id int) (Question, error)
	GetAllQuestion() ([]Question, error)
	UpdateQuestion(id int, question Question) (Question, error)
	DeleteQuestion(id int) (Question, error)
}
