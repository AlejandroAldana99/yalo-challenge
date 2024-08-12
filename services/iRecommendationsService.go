package services

import "github.com/AlejandroAldana99/yalo-challenge/models"

type IRecomendationService interface {
	GetRecomendationsByUserID(userID string) (models.Recommendation, error)
}
