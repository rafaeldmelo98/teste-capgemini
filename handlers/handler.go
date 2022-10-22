package handlers

import (
	"net/http"
	"teste-capgemini/models"

	"github.com/labstack/echo/v4"
)

type jsonObj map[string]interface{}

func CheckSequence(c echo.Context) error {
	var sequence models.Sequence
	err := c.Bind(&sequence)
	if err != nil {
		return c.JSON(http.StatusBadRequest, jsonObj{
			"error": "Error trying to decode json",
		})
	}

	return nil
}

func Stats(c echo.Context) error {
	return nil
}
