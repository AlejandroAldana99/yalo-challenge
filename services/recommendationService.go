package services

import (
	"github.com/AlejandroAldana99/yalo-challenge/errors"
	"github.com/AlejandroAldana99/yalo-challenge/libs/logger"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/repositories"
	"github.com/AlejandroAldana99/yalo-challenge/utils"
)

type RecommendationService struct {
	Repository repositories.IInteractionsRepository
}

func (service RecommendationService) GetRecommendationsByUserID(userID string) (models.Recommendation, error) {
	interactions, err := service.Repository.GetInteractionsByUserID(userID)
	if err != nil {
		logger.Error("services", "GetRecomendationsByUserID", err.Error())
		return models.Recommendation{}, errors.HandleServiceError(err)
	}

	recomendations := utils.RankTopProducts(interactions)

	response := models.Recommendation{
		UserID:   userID,
		Products: recomendations,
	}

	return response, nil
}
