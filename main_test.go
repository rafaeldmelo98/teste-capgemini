package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"teste-capgemini/handlers"
	"teste-capgemini/services"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMapMatrixSequence(t *testing.T) {
	listReceived := []string{"BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD"}
	letter := "B"
	matrix := services.MapSequence(letter, listReceived)
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

func TestValidSequence(t *testing.T) {
	validResponse := `{"is_valid":true}
`
	listSent := `{"letters":["DUHBHB","DUBUHD","UBUUHU","BHBDHH","DDDDUB","UDBDUH"]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/sequence", strings.NewReader(listSent))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handlers.CheckSequence(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, validResponse, rec.Body.String())
	}
}

func TestInvalidSequence(t *testing.T) {
	validResponse := `{"is_valid":false}
`
	listSent := `{"letters":["DUHBHB","DUBUUD","UBUUHU","HHBDHH","DHDDUB","UDBDUH"]}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/sequence", strings.NewReader(listSent))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handlers.CheckSequence(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, validResponse, rec.Body.String())
	}
}
