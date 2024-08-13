package services

import "github.com/AlejandroAldana99/yalo-challenge/models"

type IRecommendationService interface {
	GetRecommendationsByUserID(userID string) (models.Recommendation, error)
}
