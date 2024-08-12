package storage

import (
	"sync"
	"time"

	"github.com/AlejandroAldana99/yalo-challenge/models"
)

type InMemoryStore struct {
	Mu   sync.RWMutex
	Data map[string][]models.UserInteraction
}

func (store *InMemoryStore) AddInteraction(userID string, interactions []models.UserInteraction) error {
	store.Mu.Lock()
	defer store.Mu.Unlock()
	store.Data[userID] = append(store.Data[userID], interactions...)
	return nil
}

func (store *InMemoryStore) GetInteractions(userID string, since time.Time) ([]models.UserInteraction, error) {
	store.Mu.RLock()
	defer store.Mu.RUnlock()
	interactions := []models.UserInteraction{}
	for _, interaction := range store.Data[userID] {
		if interaction.Timestamp.After(since) {
			interactions = append(interactions, interaction)
		}
	}
	return interactions, nil
}
