package services

import (
	"github.com/AlejandroAldana99/yalo-challenge/errors"
	"github.com/AlejandroAldana99/yalo-challenge/libs/logger"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/repositories"
)

type InteractionsService struct {
	Repository repositories.IRecomendationRepository
}

func (service InteractionsService) CollectUserInteractions(interactions []models.UserInteraction) error {
	err := service.Repository.CollectUserInteraction(interactions)
	if err != nil {
		logger.Error("services", "CollectUserInteractions", err.Error())
		return errors.HandleServiceError(err)
	}

	return nil
}
