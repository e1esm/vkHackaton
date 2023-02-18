package db

import (
	"errors"
	"github.com/labstack/gommon/log"
	"vkHackaton/internal/cfg"
	"vkHackaton/internal/models"
)

func SaveEvent(event models.Event) error {
	query := "INSERT INTO nft_tickets_app.events (Id, Title, Description) VALUES (?, ?, ?)"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		return errors.New("couldn't prepare statement to save event")
	}
	_, err = stmt.Query(event.ID, event.Title, event.Description)
	if err != nil {
		return errors.New("Couldn't save event after preparing statement")
	}
	return nil
}

func QueryForAllEvents() ([]models.Event, error) {
	queriedEvents := make([]models.Event, 0)
	query := "SELECT * FROM nft_tickets_app.events"
	rows, err := cfg.CFG.DB.Query(query)
	if err != nil {
		log.Debug("Couldn't query for all events", err)
	}
	for rows.Next() {
		var currentModel models.Event = models.Event{}

		if er := rows.Scan(&currentModel.ID, &currentModel.Title, &currentModel.Description); er != nil {
			return queriedEvents, er
		}
		queriedEvents = append(queriedEvents, currentModel)
	}
	return queriedEvents, nil
}

func DeleteEvent(title string) error {
	query := "DELETE FROM nft_tickets_app.events WHERE (`title` = '?')"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		return err
	}
	_ = stmt.QueryRow(title)
	return nil
}

func EventQuery(title string) models.Event {
	var event models.Event = models.Event{}
	query := "SELECT * FROM nft_tickets_app.events WHERE nft_tickets_app.events.title = ?"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		log.Debugf("Couldn't fetch an entry from DB with this id: %s", err)
	}
	row := stmt.QueryRow(title)
	if err := row.Scan(&event.ID, &event.Title, &event.Description); err != nil {
		return event
	}
	return event
}
