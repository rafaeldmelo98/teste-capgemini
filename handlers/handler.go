package handlers

import (
	"database/sql"
	"fmt"
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
			"error":   "Error trying to decode json",
			"message": err.Error(),
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

	quantitySequenceFound := foundSequenceB + foundSequenceU + foundSequenceD +
		foundSequenceH
	rowsMatrix, columnsMatrix := services.GetMatrixSize(matrixB)
	quantityElementsMatrix := rowsMatrix * columnsMatrix
	numberElementsValid := quantitySequenceFound * 4
	numberElementsInvalid := quantityElementsMatrix - numberElementsValid
	rateNumberElementValid := float64(numberElementsValid) / float64(quantityElementsMatrix)

	query := fmt.Sprintf(`INSERT INTO sequences(quantity_valid_sequence,quantity_invalid_sequence,rate_valid_sequence) VALUES(%d, %d,%.1f)`,
		numberElementsValid, numberElementsInvalid, rateNumberElementValid)
	statment, err := handler.DB.Prepare(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonObj{
			"error":   "There is a problem saving data",
			"message": err.Error(),
		})
	}
	_, err = statment.Exec()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonObj{
			"error":   "There is a problem saving data",
			"message": err.Error(),
		})
	}

	if quantitySequenceFound >= 2 {
		return c.JSON(http.StatusOK, jsonObj{
			"is_valid": true,
		})
	}

	return c.JSON(http.StatusOK, jsonObj{
		"is_valid": false,
	})
}

func (handler *Handler) Stats(c echo.Context) error {
	rowsSelected, err := handler.DB.Query("SELECT quantity_valid_sequence,quantity_invalid_sequence,rate_valid_sequence from sequences")
	var tableResult []models.SequenceTable
	var record models.SequenceTable
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonObj{
			"error":   "Error trying to select data",
			"message": err.Error(),
		})
	}
	defer rowsSelected.Close()
	for rowsSelected.Next() {
		rowsSelected.Scan(&record.QuantityValidSequence, &record.QuantityInvalidSequence, &record.RateValidSequence)
		tableResult = append(tableResult, record)
	}
	return c.JSON(http.StatusOK, tableResult)
}
