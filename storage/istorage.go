package storage

import (
	"time"

	"github.com/AlejandroAldana99/yalo-challenge/models"
)

type IStoreRepository interface {
	AddInteraction(userID string, interactions []models.UserInteraction) error
	GetInteractions(userID string, since time.Time) ([]models.UserInteraction, error)
}
