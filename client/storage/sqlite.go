package storage

import (
	"database/sql"
	"fmt"

	"proyecto-chat/domain"

	_ "modernc.org/sqlite"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(archivo string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", archivo)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
        id        TEXT PRIMARY KEY,
        author    TEXT NOT NULL,
        text      TEXT NOT NULL,
        timestamp TEXT NOT NULL
    )`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) Save(message domain.Message) error {
	_, err := s.db.Exec(
		"INSERT INTO messages (id, author, text, timestamp) VALUES (?, ?, ?, ?)",
		message.ID, message.Author, message.Text, message.Timestamp,
	)
	return err
}

func (s *SQLiteStorage) FindAll() ([]domain.Message, error) {
	rows, err := s.db.Query("SELECT id, author, text, timestamp FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mensajes []domain.Message
	for rows.Next() {
		var msg domain.Message
		err := rows.Scan(&msg.ID, &msg.Author, &msg.Text, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		mensajes = append(mensajes, msg)
	}
	return mensajes, nil
}

func (s *SQLiteStorage) Delete(id string) error {
	result, err := s.db.Exec("DELETE FROM messages WHERE id = ?", id)
	if err != nil {
		return err
	}
	filas, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if filas == 0 {
		return fmt.Errorf("mensaje no encontrado")
	}
	return nil
}
