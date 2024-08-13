package repositories

import (
	"time"

	"github.com/AlejandroAldana99/yalo-challenge/config"
	"github.com/AlejandroAldana99/yalo-challenge/errors"
	"github.com/AlejandroAldana99/yalo-challenge/libs/logger"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/storage"
)

type InteractionsRepository struct {
	Config  config.Configuration
	Storage storage.IStoreRepository
}

func (repo InteractionsRepository) GetInteractionsByUserID(userID string) ([]models.UserInteraction, error) {
	t := time.Now()
	since := t.AddDate(0, 0, -7)

	recomendations, err := repo.Storage.GetInteractions(userID, since)
	if err != nil {
		logger.Error("repositories", "GetInteractionsByUserID", err.Error())
		return []models.UserInteraction{}, errors.HandleServiceError(err)
	}

	logger.Performance("repository", "GetInteractionsByUserID", t)
	return recomendations, nil
}

func (repo InteractionsRepository) CollectUserInteraction(user []models.UserInteraction) error {
	t := time.Now()
	err := repo.Storage.AddInteraction(user[0].UserID, user)
	if err != nil {
		logger.Error("repositories", "CollectUserInteraction", err.Error())
		return errors.HandleServiceError(err)
	}

	logger.Performance("repository", "CollectUserInteraction", t)

	return nil
}
