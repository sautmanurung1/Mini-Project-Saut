package discussions_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"Tugas-Mini-Project/testing/repo_test/discussions"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateDiscussionsRepository(t *testing.T) {
	discussionsRepo := new(discussions.DiscussionsRepository)
	discussionsData := entities.Discussions{
		UserId:     1,
		QuestionId: 1,
		AnswareId:  1,
	}

	t.Run("Success", func(t *testing.T) {
		discussionsRepo.On("CreateDiscussions", mock.Anything).Return(nil).Once()

		discussionsRepo := domains.DiscussionsRepository(discussionsRepo)
		err := discussionsRepo.CreateDiscussions(discussionsData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussionsRepo.On("CreateDiscussions", mock.Anything).Return(errors.New("Error To Create Discussions")).Once()

		discussionsRepo := domains.DiscussionsRepository(discussionsRepo)
		err := discussionsRepo.CreateDiscussions(discussionsData)
		assert.Error(t, err)
	})
}

func TestGetAllDiscussionsRepository(t *testing.T) {
	discussionsRepo := new(discussions.DiscussionsRepository)
	discussionsData := []entities.Discussions{
		{
			UserId:       1,
			Name:         "User Testing",
			QuestionId:   1,
			QuestionUser: "Question Testing",
			AnswareId:    1,
			AnswereUser:  "Answare Testing",
		},
	}

	t.Run("Success", func(t *testing.T) {
		discussionsRepo.On("GetAllDiscussions").Return(discussionsData, nil).Once()

		discussRepo := domains.DiscussionsRepository(discussionsRepo)
		getAllDiscuss, err := discussRepo.GetAllDiscussions()

		assert.Equal(t, getAllDiscuss, discussionsData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussionsRepo.On("GetAllDiscussions").Return(discussionsData, errors.New("Error To Create Discussions")).Once()

		discussRepo := domains.DiscussionsRepository(discussionsRepo)
		getAllDiscuss, err := discussRepo.GetAllDiscussions()

		assert.Equal(t, getAllDiscuss, discussionsData)
		assert.Error(t, err)
	})
}
