package handlers

import (
	"database/sql"
	"net/http"
	"teste-capgemini/models"
	"teste-capgemini/services"

	"github.com/labstack/echo/v4"
)

type jsonObj map[string]interface{}

type Handler struct {
	DB *sql.DB
}

func (handler *Handler) CheckSequence(c echo.Context) error {
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

	foundSequenceB := services.FindValidSequence(matrixB)
	foundSequenceU := services.FindValidSequence(matrixU)
	foundSequenceD := services.FindValidSequence(matrixD)
	foundSequenceH := services.FindValidSequence(matrixH)

	rowsMatrix, columnsMatrix := services.GetMatrixSize(matrixB)
	quantityElementsMatrix := rowsMatrix * columnsMatrix
	quantitySequenceFounded := foundSequenceB + foundSequenceU + foundSequenceD +
		foundSequenceH
	numberElementsValid := quantitySequenceFounded * 4
	numberElementsInvalid := quantityElementsMatrix - numberElementsValid
	rateNumberElementValid := float64(numberElementsValid) / float64(quantityElementsMatrix)

	if quantitySequenceFounded >= 2 {
		return c.JSON(http.StatusOK, jsonObj{
			"is_valid":  true,
			"sequences": quantitySequenceFounded,
			"n_valid":   numberElementsValid,
			"n_invalid": numberElementsInvalid,
			"rate":      rateNumberElementValid,
		})
	}

	return c.JSON(http.StatusOK, jsonObj{
		"is_valid": false,
	})
}

func (handler *Handler) Stats(c echo.Context) error {
	return nil
}
