package discussions

import (
	"Tugas-Mini-Project/domains"
	"Tugas-Mini-Project/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateDiscussionsService(t *testing.T) {
	discussService := new(DiscussionsService)
	discussionsData := entities.Discussions{
		UserId:     1,
		QuestionId: 1,
		AnswareId:  1,
	}

	t.Run("Success", func(t *testing.T) {
		discussService.On("CreateDiscussionsService", mock.Anything).Return(nil).Once()

		discussService := domains.DiscussionsService(discussService)
		err := discussService.CreateDiscussionsService(discussionsData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		discussService.On("CreateDiscussionsService", mock.Anything).Return(errors.New("Error To Make Unit Testing")).Once()

		discussService := domains.DiscussionsService(discussService)
		err := discussService.CreateDiscussionsService(discussionsData)

		assert.Error(t, err)
	})
}

func TestGetAllDiscussionsService(t *testing.T) {
	discussService := new(DiscussionsService)
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
