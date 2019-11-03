package model

import "time"

// Note struct contain information about a note
type Note struct {
	ID       uint64    `json:"note_id"`
	Title    string    `json:"note_name"`
	Detail   string    `json:"detail"`
	LastEdit time.Time `json:"last_edit"`
}
