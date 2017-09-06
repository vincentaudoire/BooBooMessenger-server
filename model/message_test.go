package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMarshalJSONWithNullPrinted(t *testing.T) {
	message := Message{
		ID:       "123",
		Message:  "Hello world",
		Received: time.Now().UTC(),
	}

	rawMessage, err := message.MarshalJSON()

	if err != nil {
		t.Error(err)
	}

	var m map[string]string
	err = json.Unmarshal(rawMessage, &m)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "", m["printed"])
}

func TestMarshalJSONWithPrinted(t *testing.T) {

	printedTime := time.Now().UTC()

	message := Message{
		ID:       "123",
		Message:  "Hello world",
		Received: time.Now(),
		Printed: mysql.NullTime{
			Valid: true,
			Time:  printedTime},
	}

	rawMessage, err := message.MarshalJSON()

	if err != nil {
		t.Error(err)
	}

	var m map[string]interface{}
	json.Unmarshal(rawMessage, &m)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, printedTime.Format(time.RFC3339Nano), m["printed"])
}
