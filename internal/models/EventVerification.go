package models

type EventVerification struct {
	Event       Event
	IsWhiteList bool
	IsAdmin     bool
}
