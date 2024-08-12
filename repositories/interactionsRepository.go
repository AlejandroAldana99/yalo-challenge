package repositories

import (
	"time"

	"github.com/AlejandroAldana99/yalo-challenge/config"
	"github.com/AlejandroAldana99/yalo-challenge/errors"
	"github.com/AlejandroAldana99/yalo-challenge/libs/logger"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/storage"
)

type RecomendationRepository struct {
	Config  config.Configuration
	Storage storage.IStoreRepository
}

func (repo RecomendationRepository) GetInteractionsByUserID(userID string) ([]models.UserInteraction, error) {
	t := time.Now()
	since := t.AddDate(0, 0, -7)

	recomendations, err := repo.Storage.GetInteractions(userID, since)
	if err != nil {
		logger.Error("repositories", "GetRecomendations", err.Error())
		return []models.UserInteraction{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetRecomendations", t)
	return recomendations, nil
}

func (repo RecomendationRepository) CollectUserInteraction(user []models.UserInteraction) error {
	t := time.Now()
	err := repo.Storage.AddInteraction(user[0].UserID, user)
	if err != nil {
		logger.Error("repositories", "CollectUserInteraction", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CollectUserInteraction", t)

	return nil
}
