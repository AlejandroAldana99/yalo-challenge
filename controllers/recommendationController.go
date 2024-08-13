package controllers

import (
	"net/http"
	"strings"

	"github.com/AlejandroAldana99/yalo-challenge/services"
	"github.com/labstack/echo/v4"
)

type RecommendationController struct {
	Service services.IRecommendationService
}

func (controller RecommendationController) GetRecomendations(c echo.Context) error {
	userID := strings.ToLower(c.Param("user_id"))
	data, err := controller.Service.GetRecommendationsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, data)
}
