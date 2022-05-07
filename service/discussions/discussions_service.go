package discussions

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/infrastructure/database"
)

type svcDiscussions struct {
	c    database.Config
	repo domains.DiscussionsRepository
}

func NewDiscussionsService(repo domains.DiscussionsRepository, c database.Config) domains.DiscussionsService {
	return &svcDiscussions{
		c:    c,
		repo: repo,
	}
}

func (s *svcDiscussions) CreateDiscussionsService(discussions entities.Discussions) (string, error) {
	return "You can Create some Discussions", s.repo.CreateDiscussions(discussions)
}

func (s *svcDiscussions) GetAllDiscussionsService() (discussions []entities.Discussions, err error) {
	return s.repo.GetAllDiscussions()
}
