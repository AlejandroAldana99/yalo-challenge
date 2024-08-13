package repositories

import "github.com/AlejandroAldana99/yalo-challenge/models"

type IInteractionsRepository interface {
	GetInteractionsByUserID(userID string) ([]models.UserInteraction, error)
	CollectUserInteraction(user []models.UserInteraction) error
}
