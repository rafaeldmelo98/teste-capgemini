package handlers

import (
	"net/http"
	"teste-capgemini/models"
	"teste-capgemini/services"

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

	matrixB := services.MapSequence("B", sequence.SequenceList)
	matrixU := services.MapSequence("U", sequence.SequenceList)
	matrixD := services.MapSequence("D", sequence.SequenceList)
	matrixH := services.MapSequence("H", sequence.SequenceList)

	foundedSequenceB := services.FindValidSequence(matrixB)
	foundedSequenceU := services.FindValidSequence(matrixU)
	foundedSequenceD := services.FindValidSequence(matrixD)
	foundedSequenceH := services.FindValidSequence(matrixH)

	quantitySequenceFounded := foundedSequenceB + foundedSequenceU + foundedSequenceD +
		foundedSequenceH

	if quantitySequenceFounded >= 2 {
		return c.JSON(http.StatusOK, jsonObj{
			"is_valid": true,
		})
	}

	return c.JSON(http.StatusOK, jsonObj{
		"is_valid": false,
	})
}

func Stats(c echo.Context) error {
	return nil
}
