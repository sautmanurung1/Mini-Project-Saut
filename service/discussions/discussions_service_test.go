package discussions_test

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/domains/mocks"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateDiscussionsService(t *testing.T) {
	discussService := new(mocks.DiscussionsService)
	discussionsData := entities.Discussions{
		UserId:     1,
		QuestionId: 1,
		AnswareId:  1,
	}

	t.Run("Success", func(t *testing.T) {
		discussService.On("CreateDiscussionsService", mock.Anything).Return("Success Create Discussions", nil).Once()

		discussService := domains.DiscussionsService(discussService)
		discuss, err := discussService.CreateDiscussionsService(discussionsData)

		assert.Equal(t, discuss, "Success Create Discussions")
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussService.On("CreateDiscussionsService", mock.Anything).Return("Failed Create Discussions", errors.New("Error To Make Unit Testing")).Once()

		discussService := domains.DiscussionsService(discussService)
		discuss, err := discussService.CreateDiscussionsService(discussionsData)

		assert.Equal(t, discuss, "Failed Create Discussions")
		assert.Error(t, err)
	})
}

func TestGetAllQuestionService(t *testing.T) {
	discussService := new(mocks.DiscussionsService)
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
		discussService.On("GetAllDiscussionsService").Return(discussionsData, nil).Once()

		discussService := domains.DiscussionsService(discussService)
		getAllDiscussions, err := discussService.GetAllDiscussionsService()

		assert.Equal(t, getAllDiscussions, discussionsData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussService.On("GetAllDiscussionsService").Return(discussionsData, errors.New("Error To Unit Testing")).Once()

		discussService := domains.DiscussionsService(discussService)
		getAllDiscussions, err := discussService.GetAllDiscussionsService()

		assert.Equal(t, getAllDiscussions, discussionsData)
		assert.Error(t, err)
	})
}
