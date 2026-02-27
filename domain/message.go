package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	Author    string
	Text      string
	Timestamp time.Time
}

func NewMessage(author, text string) (Message, error) {
	if author == "" {
		return Message{}, fmt.Errorf("el autor no puede estar vacío")
	}
	if text == "" {
		return Message{}, fmt.Errorf("el texto no puede estar vacío")
	}

	return Message{
		ID:        uuid.NewString(),
		Author:    author,
		Text:      text,
		Timestamp: time.Now(),
	}, nil
}
