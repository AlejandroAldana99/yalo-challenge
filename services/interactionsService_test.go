package services

import (
	"testing"

	"github.com/AlejandroAldana99/yalo-challenge/mocks"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestCollectUserInteractions(t *testing.T) {
	mockRepo := new(mocks.IRecomendationRepository)
	service := InteractionsService{
		Repository: mockRepo,
	}

	tests := []struct {
		name          string
		interactions  []models.UserInteraction
		mockError     error
		expectedError error
	}{
		{
			name: "Successful collection of interactions",
			interactions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440000", Duration: 10},
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440001", Duration: 20},
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "No interactions provided",
			interactions:  []models.UserInteraction{},
			mockError:     nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			mockRepo.On("CollectUserInteraction", tt.interactions).Return(tt.mockError)

			// Execute the service function
			err := service.CollectUserInteractions(tt.interactions)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure that the expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}
