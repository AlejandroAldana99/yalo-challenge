package routes

import (
	"net/http"

	"github.com/AlejandroAldana99/yalo-challenge/controllers"
	"github.com/AlejandroAldana99/yalo-challenge/libs/logger"
	"github.com/AlejandroAldana99/yalo-challenge/middleware"
	"github.com/AlejandroAldana99/yalo-challenge/server/di"
	"github.com/labstack/echo/v4"
)

// Route represents the route structure for the service
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// Routes represents a route collection
type Routes []Route

// ServiceRoutes is a route collection for service handling
var ServiceRoutes Routes

func init() {
	controllersProvider := di.BuildContainer()
	err := controllersProvider.Invoke(func(
		healthController *controllers.HealthController,
		interactionsController *controllers.InteractionsController,
		recommendationController *controllers.RecommendationController,
	) {
		ServiceRoutes = Routes{
			Route{
				Method:      http.MethodGet,
				Pattern:     "/health",
				HandlerFunc: healthController.HealthCheck,
				Name:        "HealthCheck",
			},
			Route{
				Method:      http.MethodGet,
				Pattern:     "/recommendations/:user_id",
				HandlerFunc: middleware.ValidatorParams(recommendationController.GetRecomendations),
				Name:        "Recommendations",
			},
			Route{
				Method:      http.MethodPost,
				Pattern:     "interactions/collect",
				HandlerFunc: middleware.ValidateBody(interactionsController.CollectUserInteraction),
				Name:        "CollectUser",
			},
		}
	})

	if err != nil {
		logger.Error("routes", "init", err.Error())
	}
}
