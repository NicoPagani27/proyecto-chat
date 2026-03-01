package storage

import (
	"fmt"
	"sync"

	"github.com/NicoPagani27/proyecto-chat/domain"
)

type MemoryStorage struct {
	mu       sync.Mutex
	mensajes []domain.Message
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Save(message domain.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mensajes = append(m.mensajes, message)
	return nil
}

func (m *MemoryStorage) FindAll() ([]domain.Message, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.mensajes, nil
}

func (m *MemoryStorage) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, msg := range m.mensajes {
		if msg.ID == id {
			m.mensajes = append(m.mensajes[:i], m.mensajes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("mensaje no encontrado")
}
