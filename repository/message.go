package repository

import (
	"BooBooMessenger-server/model"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// MessageRepository .
type MessageRepository struct {
	db *sql.DB
}

// New .
func New(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// GetAllMessages .
func (r *MessageRepository) GetAllMessages() ([]model.Message, error) {
	statement, err := r.db.Prepare(`SELECT id, message, received
		FROM message
		WHERE printed IS NULL
		ORDER BY received DESC`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []model.Message{}

	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.ID, &m.Message, &m.Received)

		if err != nil {
			return nil, err
		}

		messages = append(messages, m)
	}

	return messages, nil
}

// MarkMessageAsRead .
func (r *MessageRepository) MarkMessageAsRead(messageID string) error {
	statement := "UPDATE message SET printed=NOW() WHERE id=$1"
	request, err := r.db.Exec(statement, messageID)

	if err != nil {
		return err
	}

	n, err := request.RowsAffected()
	if err != nil {
		return err
	}

	if n != 1 {
		return fmt.Errorf("Couldn't find message with id: %s", messageID)
	}

	return nil
}

// SaveMessage .
func (r *MessageRepository) SaveMessage(message *model.Message) (*model.Message, error) {

	messageToSave := &model.Message{
		ID:       uuid.New().String(),
		Message:  message.Message,
		Received: time.Now(),
	}

	statement := `INSERT INTO message (id, message, received) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(statement, messageToSave.ID, messageToSave.Message, messageToSave.Received)
	if err != nil {
		return nil, err
	}

	return messageToSave, nil
}
