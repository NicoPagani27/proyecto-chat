package domain

import "fmt"

// --- Enviar Mensaje ---
type SendMessageUseCase struct {
	repo MessageRepository
}

func NewSendMessageUseCase(repo MessageRepository) *SendMessageUseCase {
	return &SendMessageUseCase{repo: repo}
}

func (uc *SendMessageUseCase) Execute(author, text string) (Message, error) {
	mensaje, err := NewMessage(author, text)
	if err != nil {
		return Message{}, err
	}
	err = uc.repo.Save(mensaje)
	return mensaje, err
}

// --- Lista de Mensajes ---
type ListMessagesUseCase struct {
	repo MessageRepository
}

func NewListMessagesUseCase(repo MessageRepository) *ListMessagesUseCase {
	return &ListMessagesUseCase{repo: repo}
}

func (uc *ListMessagesUseCase) Execute() ([]Message, error) {
	return uc.repo.FindAll()
}

// --- Mensaje Eliminado ---
type DeleteMessageUseCase struct {
	repo MessageRepository
}

func NewDeleteMessageUseCase(repo MessageRepository) *DeleteMessageUseCase {
	return &DeleteMessageUseCase{repo: repo}
}

func (uc *DeleteMessageUseCase) Execute(id string) error {
	if id == "" {
		return fmt.Errorf("el id no puede estar vacío")
	}
	return uc.repo.Delete(id)
}
