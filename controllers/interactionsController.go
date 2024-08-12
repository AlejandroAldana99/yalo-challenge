package controllers

import (
	"net/http"

	"github.com/AlejandroAldana99/yalo-challenge/models"
	"github.com/AlejandroAldana99/yalo-challenge/services"
	"github.com/labstack/echo/v4"
)

type InteractionsController struct {
	Service services.IInteractionsService
}

func (controller InteractionsController) CollectUserInteraction(c echo.Context) error {
	dto := c.Get("dto").([]models.UserInteraction)
	err := controller.Service.CollectUserInteractions(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}
