package services

import (
	"errors"
	"testing"

	"github.com/AlejandroAldana99/yalo-challenge/mocks"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestGetRecomendationsByUserID(t *testing.T) {
	mockRepo := new(mocks.IRecomendationRepository)
	service := RecomendationService{
		Repository: mockRepo,
	}

	tests := []struct {
		name             string
		userID           string
		mockInteractions []models.UserInteraction
		mockError        error
		expectedResult   models.Recommendation
		expectedError    error
	}{
		{
			name:   "User with interactions",
			userID: "user1",
			mockInteractions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440000", Duration: 10},
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440001", Duration: 20},
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440002", Duration: 30},
			},
			mockError: nil,
			expectedResult: models.Recommendation{
				UserID:   "user1",
				Products: []string{"550e8400-e29b-41d4-a716-446655440002", "550e8400-e29b-41d4-a716-446655440001", "550e8400-e29b-41d4-a716-446655440000"},
			},
			expectedError: nil,
		},
		{
			name:             "User with no interactions",
			userID:           "user2",
			mockInteractions: []models.UserInteraction{},
			mockError:        nil,
			expectedResult:   models.Recommendation{UserID: "user2", Products: []string{}},
			expectedError:    nil,
		},
		{
			name:             "Repository error",
			userID:           "user3",
			mockInteractions: nil,
			mockError:        errors.New("repository error"),
			expectedResult:   models.Recommendation{},
			expectedError:    errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			mockRepo.On("GetInteractionsByUserID", tt.userID).Return(tt.mockInteractions, tt.mockError)

			// Execute the service function
			result, err := service.GetRecomendationsByUserID(tt.userID)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}

			// Ensure that the expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}
