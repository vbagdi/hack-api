package models

type EventAttendees struct {
	Event     string   `json:"event"`
	Attendees []string `json:"attendees"`
}
