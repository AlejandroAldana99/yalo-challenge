package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlejandroAldana99/yalo-challenge/mocks"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetRecommendations(t *testing.T) {
	mockService := new(mocks.IRecommendationService)
	controller := RecommendationController{
		Service: mockService,
	}

	tests := []struct {
		name             string
		userID           string
		mockData         models.Recommendation
		mockError        error
		expectedStatus   int
		expectedResponse interface{}
	}{
		{
			name:   "Successful recommendation retrieval",
			userID: "user1",
			mockData: models.Recommendation{
				UserID:   "user1",
				Products: []string{"550e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440001"},
			},
			mockError:        nil,
			expectedStatus:   http.StatusOK,
			expectedResponse: models.Recommendation{UserID: "user1", Products: []string{"550e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440001"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup echo context
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set the user ID parameter in the context
			c.SetParamNames("user_id")
			c.SetParamValues(tt.userID)

			// Setup mock expectations
			mockService.On("GetRecommendationsByUserID", tt.userID).Return(tt.mockData, tt.mockError)

			// Execute the controller function
			err := controller.GetRecomendations(c)

			// Assertions
			if tt.mockError != nil && err != nil {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.expectedResponse.(string))
			} else {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.JSONEq(t, `{"user_id":"`+tt.userID+`","products":["550e8400-e29b-41d4-a716-446655440000","550e8400-e29b-41d4-a716-446655440001"]}`, rec.Body.String())
			}

			// Ensure that the expectations were met
			mockService.AssertExpectations(t)
		})
	}
}
