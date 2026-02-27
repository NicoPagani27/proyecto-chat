package domain

import (
	"fmt"
	"testing"
)

// --- fake repo para los tests ---
type fakeRepo struct {
	mensajes []Message
}

func (f *fakeRepo) Save(message Message) error {
	f.mensajes = append(f.mensajes, message)
	return nil
}

func (f *fakeRepo) FindAll() ([]Message, error) {
	return f.mensajes, nil
}

func (f *fakeRepo) Delete(id string) error {
	for i, msg := range f.mensajes {
		if msg.ID == id {
			f.mensajes = append(f.mensajes[:i], f.mensajes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("mensaje no encontrado")
}

// --- Tests de la entidad ---
func TestNewMessage_AutorVacio(t *testing.T) {
	_, err := NewMessage("", "Hola")
	if err == nil {
		t.Error("debería retornar error con autor vacío")
	}
}

func TestNewMessage_TextoVacio(t *testing.T) {
	_, err := NewMessage("Nico", "")
	if err == nil {
		t.Error("debería retornar error con texto vacío")
	}
}

func TestNewMessage_Valido(t *testing.T) {
	msg, err := NewMessage("Nico", "Hola")
	if err != nil {
		t.Errorf("no debería retornar error: %v", err)
	}
	if msg.ID == "" {
		t.Error("el ID no puede estar vacío")
	}
	if msg.Timestamp.IsZero() {
		t.Error("el timestamp no puede ser cero")
	}
}

// --- Tests de los use cases ---
func TestSendMessage_Valido(t *testing.T) {
	repo := &fakeRepo{}
	uc := NewSendMessageUseCase(repo)

	msg, err := uc.Execute("Nico", "Hola")
	if err != nil {
		t.Errorf("no debería retornar error: %v", err)
	}
	if len(repo.mensajes) != 1 {
		t.Error("el mensaje debería haberse guardado en el repo")
	}
	if msg.Author != "Nico" {
		t.Error("el autor no coincide")
	}
}

func TestSendMessage_DatosInvalidos(t *testing.T) {
	repo := &fakeRepo{}
	uc := NewSendMessageUseCase(repo)

	_, err := uc.Execute("", "Hola")
	if err == nil {
		t.Error("debería retornar error con autor vacío")
	}
	if len(repo.mensajes) != 0 {
		t.Error("no debería haberse guardado nada en el repo")
	}
}

func TestDeleteMessage_Existente(t *testing.T) {
	repo := &fakeRepo{}
	send := NewSendMessageUseCase(repo)
	delete := NewDeleteMessageUseCase(repo)

	msg, _ := send.Execute("Nico", "Hola")
	err := delete.Execute(msg.ID)
	if err != nil {
		t.Errorf("no debería retornar error: %v", err)
	}
	if len(repo.mensajes) != 0 {
		t.Error("el mensaje debería haberse eliminado")
	}
}

func TestDeleteMessage_Inexistente(t *testing.T) {
	repo := &fakeRepo{}
	uc := NewDeleteMessageUseCase(repo)

	err := uc.Execute("id-que-no-existe")
	if err == nil {
		t.Error("debería retornar error con ID inexistente")
	}
}
