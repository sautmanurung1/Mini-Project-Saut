package domains

import "Tugas-Mini-Project/entities"

type DiscussionsRepository interface {
	GetAllDiscussions(discussions entities.Discussions) error
}

type DiscussionsService interface {
	GetAllDiscussionsService(discussions entities.Discussions) error
}
