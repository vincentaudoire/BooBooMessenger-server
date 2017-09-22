package repository_test

import (
	"testing"
	"time"

	messageRepository "github.com/vincentaudoire/BooBooMessenger-server/repository"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllMessages(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "message", "received"}).
		AddRow(1, "title 1", time.Now()).
		AddRow(2, "title 2", time.Now())

	query := `SELECT id, message, received
		FROM message
		WHERE printed IS NULL
		ORDER BY received DESC`

	mock.
		ExpectPrepare(query).
		ExpectQuery().
		WillReturnRows(rows)

	r := messageRepository.New(db)
	messages, err := r.GetAllMessages()

	assert.NoError(t, err)
	assert.Len(t, messages, 2)
}
