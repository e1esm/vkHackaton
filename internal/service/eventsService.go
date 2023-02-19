package service

import (
	"github.com/labstack/gommon/log"
	"vkHackaton/internal/db"
	"vkHackaton/internal/models"
)

func FetchAllEvents() []models.Event {
	events, err := db.QueryForAllEvents()
	if err != nil {
		log.Fatal(err)
	}
	return events
}

func FetchEvent(title string) models.Event {
	event := db.EventQuery(title)
	return event
}

func SetCurrentEvent(event models.Event, lastOrganizedId int) error {
	err := db.SaveEvent(event, lastOrganizedId)
	return err
}

func GetOrganizerId() int {
	id := db.GetOrganizerId()
	return id
}

func Authorization(user models.User) {
	db.Authorize(user)
}

func GetWallet(userId int) string {
	wallet := db.GetWallet(userId)
	return wallet
}

func VerifyUser(event models.Event, userId int) bool {
	isWhiteListed := db.IsInWhiteList(event.ID, userId)
	return isWhiteListed
}

func IsAdmin(event models.Event, userId int) bool {
	isAdmin := db.IsAdmin(event.ID, userId)

	return isAdmin
}
