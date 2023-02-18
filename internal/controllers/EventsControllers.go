package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"vkHackaton/internal/models"
	"vkHackaton/internal/service"
)

func GetEvent(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	title := c.Param("title")

	event := service.FetchEvent(title)
	if event.Title == "" {
		return c.JSON(http.StatusOK, "Not found")
	} else {
		return c.JSON(http.StatusOK, event)
	}
}

func GetAllEvents(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	events := service.FetchAllEvents()
	return c.JSON(http.StatusOK, events)
}

func EditEvent(c echo.Context) error {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Fatal("Empty json body")
	}

	return c.JSON(http.StatusOK, "")
}

/*
func EditEvent(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Fatal("Couldn't have encoded body(PUT request)")
	}

}

*/

func SaveEvent(c echo.Context) error {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Fatal("Empty json body")
	}
	event := models.Event{ID: uuid.New(), Title: fmt.Sprintf("%s", jsonBody["title"]),
		Description: fmt.Sprintf("%s", jsonBody["description"])}
	service.SetCurrentEvent(event)

	c.Response().Header().Set("Content-Type", "text/plain")
	return c.JSON(http.StatusOK, event)
}
