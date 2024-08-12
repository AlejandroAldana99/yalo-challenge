package services

import "github.com/AlejandroAldana99/yalo-challenge/models"

type IHealthService interface {
	CheckPod(chanHealth chan models.HealthComponentDetail)
}
