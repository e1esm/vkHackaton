package service

import (
	"vkHackaton/internal/db"
	"vkHackaton/internal/models"
)

func FetchAllEvents() []models.Event {
	events, err := db.QueryForAllEvents()
	if err != nil {
		return nil
	}
	return events
}

func FetchEvent(title string) models.Event {
	event := db.EventQuery(title)
	return event
}

func SetCurrentEvent(event models.Event) error {
	err := db.SaveEvent(event)
	return err
}
