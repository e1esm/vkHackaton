package models

import "github.com/google/uuid"

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Image       string    `json:"image"`
	Format      string    `json:"format"`
	Link        string    `json:"link"`
	Tag         string    `json:"tag"`
	Owner       Owner     `json:"owner"`
}
