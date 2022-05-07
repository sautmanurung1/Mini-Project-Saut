package domains

import "Tugas-Mini-Project/entities"

type DiscussionsRepository interface {
	GetAllDiscussions() ([]entities.Discussions, error)
}

type DiscussionsService interface {
	GetAllDiscussionsService() (discuss []entities.Discussions, err error)
}
