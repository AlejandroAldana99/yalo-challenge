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

func TestCollectUserInteraction(t *testing.T) {
	mockService := new(mocks.IInteractionsService)
	controller := InteractionsController{
		Service: mockService,
	}

	tests := []struct {
		name             string
		dto              []models.UserInteraction
		mockError        error
		expectedStatus   int
		expectedResponse interface{}
	}{
		{
			name: "Successful collection",
			dto: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440000", Duration: 10},
				{UserID: "user1", ProductSKU: "550e8400-e29b-41d4-a716-446655440001", Duration: 20},
			},
			mockError:        nil,
			expectedStatus:   http.StatusOK,
			expectedResponse: nil,
		},
		{
			name:             "No interactions provided",
			dto:              []models.UserInteraction{},
			mockError:        nil,
			expectedStatus:   http.StatusOK,
			expectedResponse: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup echo context
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set the DTO in the context
			c.Set("dto", tt.dto)

			// Setup mock expectations
			mockService.On("CollectUserInteractions", tt.dto).Return(tt.mockError)

			// Execute the controller function
			err := controller.CollectUserInteraction(c)

			// Assertions
			if tt.mockError != nil && err != nil {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.expectedResponse.(string))
			} else {
				assert.Equal(t, tt.expectedStatus, rec.Code)
			}

			// Ensure that the expectations were met
			mockService.AssertExpectations(t)
		})
	}
}
