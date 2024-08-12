package services

import "github.com/AlejandroAldana99/yalo-challenge/models"

type IInteractionsService interface {
	CollectUserInteractions(interactions []models.UserInteraction) error
}
