package db

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"vkHackaton/internal/cfg"
	"vkHackaton/internal/models"
)

func SaveEvent(event models.Event, lastOrganizerId int) error {
	query := "INSERT INTO nft_tickets_app.events (Id, Title, Description, Date, Platform, Image, Link, Org_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		return errors.New("couldn't prepare statement to save event")
	}
	_, err = stmt.Query(event.ID, event.Title, event.Description, event.Date, event.Format, event.Image, event.Link, lastOrganizerId)
	if err != nil {
		return errors.New("Couldn't save event after preparing statement")
	}

	query = "INSERT INTO nft_tickets_app.organizers (Id, Name, Address) VALUES (?, ?, ?)"

	stmt, err = cfg.CFG.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(lastOrganizerId, event.Owner.Name, event.Owner.Link)

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

func GetOrganizerId() int {
	var id int
	query := "SELECT COUNT(*) FROM nft_tickets_app.organizers"
	stmt := cfg.CFG.DB.QueryRow(query)
	if err := stmt.Scan(&id); err != nil {
		log.Fatalf("Coulnd't assign amount of rows: %s", err)
	}
	return id
}

func Authorize(user models.User) {
	query := "INSERT INTO nft_tickets_app.users (Id, Public_Hash) VALUES (?, ?)"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	stmt.Query(user.UserId, user.Wallet)
}

func GetWallet(id int) string {
	var wallet string
	query := "SELECT Public_Hash from nft_tickets_app.users where Id = ?"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	row := stmt.QueryRow(id)
	if row.Err() == sql.ErrNoRows {

		return ""
	} else {
		row.Scan(&wallet)
	}
	return wallet
}

func IsInWhiteList(eventId uuid.UUID, id int) bool {

	query := "SELECT COUNT(*) FROM nft_tickets_app.white_list WHERE event_id = ? AND user_id = ?"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	var amount int
	row := stmt.QueryRow(eventId, id)
	err = row.Scan(&amount)
	if err != nil {
		log.Fatal(err.Error())
	}
	if amount == 0 {
		return false
	} else {
		return true
	}
}

func IsAdmin(eventId uuid.UUID, userId int) bool {
	query := "SELECT COUNT(*) FROM nft_tickets_app.events WHERE Org_id = ? AND id = ?"
	stmt, err := cfg.CFG.DB.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	var amountFound int

	row := stmt.QueryRow(userId, eventId)

	row.Scan(&amountFound)
	if amountFound == 0 {
		return false
	} else {
		return true
	}

}
