package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/AlejandroAldana99/yalo-challenge/models"
)

func TestRankTopProducts(t *testing.T) {
	tests := []struct {
		name         string
		interactions []models.UserInteraction
		expectedTop3 []string
	}{
		{
			name:         "No interactions",
			interactions: []models.UserInteraction{},
			expectedTop3: []string{},
		},
		{
			name: "Single interaction",
			interactions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "product1", Action: models.ActionTypeView, Timestamp: time.Now(), Duration: 10},
			},
			expectedTop3: []string{"product1"},
		},
		{
			name: "Three interactions, different products",
			interactions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "product1", Action: models.ActionTypeView, Timestamp: time.Now(), Duration: 10},
				{UserID: "user1", ProductSKU: "product2", Action: models.ActionTypeClick, Timestamp: time.Now(), Duration: 20},
				{UserID: "user1", ProductSKU: "product3", Action: models.ActionTypeAddToCart, Timestamp: time.Now(), Duration: 30},
			},
			expectedTop3: []string{"product3", "product2", "product1"},
		},
		{
			name: "Multiple interactions with the same product",
			interactions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "product1", Action: models.ActionTypeView, Timestamp: time.Now(), Duration: 10},
				{UserID: "user1", ProductSKU: "product1", Action: models.ActionTypeClick, Timestamp: time.Now(), Duration: 20},
				{UserID: "user1", ProductSKU: "product2", Action: models.ActionTypeAddToCart, Timestamp: time.Now(), Duration: 30},
			},
			expectedTop3: []string{"product1", "product2"},
		},
		{
			name: "More than three products",
			interactions: []models.UserInteraction{
				{UserID: "user1", ProductSKU: "product1", Action: models.ActionTypeView, Timestamp: time.Now(), Duration: 10},
				{UserID: "user1", ProductSKU: "product2", Action: models.ActionTypeClick, Timestamp: time.Now(), Duration: 20},
				{UserID: "user1", ProductSKU: "product3", Action: models.ActionTypeAddToCart, Timestamp: time.Now(), Duration: 30},
				{UserID: "user1", ProductSKU: "product4", Action: models.ActionTypeView, Timestamp: time.Now(), Duration: 40},
				{UserID: "user1", ProductSKU: "product5", Action: models.ActionTypeClick, Timestamp: time.Now(), Duration: 50},
			},
			expectedTop3: []string{"product5", "product4", "product3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RankTopProducts(tt.interactions)
			if !reflect.DeepEqual(got, tt.expectedTop3) {
				t.Errorf("RankTopProducts() = %v, want %v", got, tt.expectedTop3)
			}
		})
	}
}
