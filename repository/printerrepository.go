package repository

import (
	"database/sql"
	"fmt"
)

type PrinterMessage struct {
	Text string `json:"text"`
}

type PrinterRepository struct {
	db *sql.DB
}

func NewPrinterRepository(db *sql.DB) PrinterRepository {
	return PrinterRepository{
		db: db,
	}
}

func (r *PrinterRepository) RegisterNewPrinter(uuid string, name string) error {
	statement := `INSERT INTO printer (uuid, name) VALUES ($1, $2)`
	_, err := r.db.Exec(statement, uuid, name)

	if err != nil {
		return err
	}

	return nil
}

func (r *PrinterRepository) GetAllPendingPrintingMessages(printerID string) ([]PrinterMessage, error) {

	q, err := r.db.Prepare(`SELECT id, text
  FROM Message
  WHERE printer_id=$1`)
	defer q.Close()

	if err != nil {
		return nil, err
	}

	rows, err := q.Query(printerID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	messages := make([]PrinterMessage, 0)
	for rows.Next() {
		var m PrinterMessage
		err = rows.Scan(&m.Text)

		if err != nil {
			return nil, err
		}

		messages = append(messages, m)
	}

	return messages, nil
}

func (r *PrinterRepository) MarkMessageAsPrinted(messageID string) error {
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
