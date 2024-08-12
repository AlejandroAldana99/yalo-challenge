package controllers

import (
	"net/http"

	"github.com/AlejandroAldana99/yalo-challenge/config"
	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/services"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
)

// HealthController :
type HealthController struct {
	Configuration config.Configuration
	MongoClient   *mongo.Client
	ServiceHealth services.IHealthService
}

// HealthCheck :
func (controller *HealthController) HealthCheck(c echo.Context) error {
	status := models.Pass.String()
	health := models.HealthStatus{
		Version:     "1.0",
		Status:      status,
		Description: "API HealthCheck",
		Details:     []models.HealthComponentDetail{},
	}

	chanPodHealth := make(chan models.HealthComponentDetail)
	defer closeChannels(chanPodHealth)
	go controller.ServiceHealth.CheckPod(chanPodHealth)
	podHealth := <-chanPodHealth

	health.Details = append(health.Details, podHealth)

	return c.JSON(http.StatusOK, health)
}

func closeChannels(channels ...chan models.HealthComponentDetail) {
	for _, item := range channels {
		close(item)
	}
}
