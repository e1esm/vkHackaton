package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"vkHackaton/internal/models"
	"vkHackaton/internal/service"
)

func GetEvent(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	title := c.Param("title")
	userId := c.Param("userId")

	event := service.FetchEvent(title)
	/*
		if event.Title == "" {
			return c.JSON(http.StatusOK, "Not found")
		} else {
			return c.JSON(http.StatusOK, event)

	*/

	userintId, _ := strconv.Atoi(userId)
	isInWhiteList := service.VerifyUser(event, userintId)
	isAdmin := service.IsAdmin(event, userintId)
	eventVerification := models.EventVerification{Event: event, IsAdmin: isAdmin, IsWhiteList: isInWhiteList}
	return c.JSON(http.StatusOK, eventVerification)

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

	lastOrganizerId := service.GetOrganizerId() + 1

	owner := jsonBody["owner"]
	nestedModel := owner.(map[string]interface{})
	event := models.Event{ID: uuid.New(), Title: fmt.Sprintf("%s", jsonBody["title"]),
		Description: fmt.Sprintf("%s", jsonBody["description"]), Date: fmt.Sprintf("%s", jsonBody["date"]),
		Image:  fmt.Sprintf("%s", jsonBody["image"]),
		Format: fmt.Sprintf("%s", jsonBody["platform"]),
		Tag:    fmt.Sprintf("%s", jsonBody["tags"]),
		Owner:  models.Owner{nestedModel["link"].(string), nestedModel["name"].(string)}}
	err = service.SetCurrentEvent(event, lastOrganizerId)
	if err != nil {
		log.Fatal(err)
	}
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.JSON(http.StatusOK, event)
}

func RegisterWaller(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	user := models.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusOK, false)
	}
	user.UserId, _ = strconv.Atoi(fmt.Sprintf("%s", jsonBody["userId"]))
	user.Wallet = fmt.Sprintf("%s", jsonBody["wallet"])
	service.Authorization(user)

	return c.JSON(http.StatusOK, true)
}

func GetWallet(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	wallet := service.GetWallet(userId)

	return c.JSON(http.StatusOK, wallet)
}
