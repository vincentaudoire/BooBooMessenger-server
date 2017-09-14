package model

import (
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Message .
type Message struct {
	ID       string         `json:"id"`
	Message  string         `json:"message"`
	Received time.Time      `json:"received"`
	Printed  mysql.NullTime `json:"printed"`
}

// MarshalJSON .
func (m *Message) MarshalJSON() ([]byte, error) {

	type Alias Message

	// We Check if the time if valid to return it
	var printed *time.Time
	if m.Printed.Valid {
		printed = &m.Printed.Time
	}

	return json.Marshal(&struct {
		Printed *time.Time `json:"printed"`
		*Alias
	}{
		Printed: printed,
		Alias:   (*Alias)(m),
	})
}
