package di

import (
	"sync"

	"github.com/AlejandroAldana99/yalo-challenge/config"
	"github.com/AlejandroAldana99/yalo-challenge/controllers"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/repositories"
	"github.com/AlejandroAldana99/yalo-challenge/services"
	"github.com/AlejandroAldana99/yalo-challenge/storage"
	"go.uber.org/dig"
)

// Storage
func NewInMemoryStore() *storage.InMemoryStore {
	return &storage.InMemoryStore{
		Mu:   sync.RWMutex{},
		Data: make(map[string][]models.UserInteraction),
	}
}

// Repositories
func newRecomendationRepository(storate *storage.InMemoryStore) *repositories.RecomendationRepository {
	return &repositories.RecomendationRepository{
		Config:  config.GetConfig(),
		Storage: storate,
	}
}

// Services
func newRecomendationService(repository *repositories.RecomendationRepository) *services.RecomendationService {
	return &services.RecomendationService{
		Repository: repository,
	}
}

func newInteractionsService(repository *repositories.RecomendationRepository) *services.InteractionsService {
	return &services.InteractionsService{
		Repository: repository,
	}
}

// Controllers
func newInteractionsController(service *services.InteractionsService) *controllers.InteractionsController {
	return &controllers.InteractionsController{
		Service: service,
	}
}

func newRecommendationController(service *services.RecomendationService) *controllers.RecommendationController {
	return &controllers.RecommendationController{
		Service: service,
	}
}

func newHealthController(service *services.HealthService) *controllers.HealthController {
	return &controllers.HealthController{
		Configuration: config.GetConfig(),
		ServiceHealth: service,
	}
}

func newHealthService() *services.HealthService {
	return &services.HealthService{}
}

// BuildContainer :
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(NewInMemoryStore)
	_ = container.Provide(newRecomendationRepository)
	_ = container.Provide(newRecomendationService)
	_ = container.Provide(newInteractionsService)
	_ = container.Provide(newRecommendationController)
	_ = container.Provide(newInteractionsController)
	_ = container.Provide(newHealthService)
	_ = container.Provide(newHealthController)

	return container
}
