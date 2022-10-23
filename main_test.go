package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"teste-capgemini/database"
	"teste-capgemini/handlers"
	"teste-capgemini/models"
	"teste-capgemini/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	setupDone = database.SetUpDatabase()
	db, _     = sql.Open("sqlite3", "sqlite-database.db")
	handler   = handlers.Handler{
		DB: db,
	}
)

func TestMapMatrixSequence(t *testing.T) {
	receivedList := []string{"BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD"}
	letter := "B"
	matrix := services.MapSequence(letter, receivedList)
	expectedMatrix := [][]bool{
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
	}

	expectedResult := fmt.Sprint(expectedMatrix)
	actualResult := fmt.Sprint(matrix)
	if actualResult != expectedResult {
		t.Errorf("Mapping matrix error, want %s, got %s", expectedResult, actualResult)
	}
}

func TestInvalidSequence(t *testing.T) {
	validResponse := `{"is_valid":false}
`
	sentList := `{"letters":["DUHBHB","DUBUUD","UBUUHU","HHBDHH","DHDDUB","UDBDUH"]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/sequence", strings.NewReader(sentList))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CheckSequence(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, validResponse, rec.Body.String())
	}
}

func TestValidSequence(t *testing.T) {
	validResponse := `{"is_valid":true}
`
	listSent := `{"letters":["DUHBHB","DUBUHD","UBUUHU","BHBDHH","DDDDUB","UDBDUH"]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/sequence", strings.NewReader(listSent))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CheckSequence(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, validResponse, rec.Body.String())
	}
}

func TestFoundRightNumbersOfSequences(t *testing.T) {
	receivedList := []string{"DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", "DDDDDB", "UDBDUH"}
	letter := "D"
	matrix := services.MapSequence(letter, receivedList)
	foundSequence := services.FindValidSequence(matrix)
	if foundSequence != 1 {
		t.Errorf("Found sequence error, want 1, got %d", foundSequence)
	}
}

func TestScanRoute(t *testing.T) {
	rowsSelected, _ := handler.DB.Query(`
	SELECT quantity_valid_sequence,quantity_invalid_sequence,rate_valid_sequence
	FROM sequences ORDER BY id DESC LIMIT 1`)
	var record models.SequenceTable
	defer rowsSelected.Close()
	for rowsSelected.Next() {
		rowsSelected.Scan(&record.QuantityValidSequence, &record.QuantityInvalidSequence, &record.RateValidSequence)
	}
	if record.QuantityValidSequence != 12 {
		t.Errorf("Select quantity valid sequence error, want 12, got %d", record.QuantityValidSequence)
	}
	if record.QuantityInvalidSequence != 24 {
		t.Errorf("Select quantity invalid sequence error, want 24, got %d", record.QuantityValidSequence)
	}
	if record.RateValidSequence != 0.3 {
		t.Errorf("Select rate valid sequence error, want 0.3, got %f", record.RateValidSequence)
	}
}
