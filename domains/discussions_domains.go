package domains

import (
	"Tugas-Mini-Project/entities"
)

type DiscussionsRepository interface {
	CreateDiscussions(discussions entities.Discussions) error
	GetAllDiscussions() (discussions []entities.Discussions, err error)
}

type DiscussionsService interface {
	CreateDiscussionsService(discussions entities.Discussions) error
	GetAllDiscussionsService() (discussions []entities.Discussions, err error)
}
