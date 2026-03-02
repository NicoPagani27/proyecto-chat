package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"proyecto-chat/domain"
)

type DiskStorage struct {
	archivo string
}

func NewDiskStorage(archivo string) *DiskStorage {
	return &DiskStorage{archivo: archivo}
}

func (d *DiskStorage) leerMensajes() ([]domain.Message, error) {
	data, err := os.ReadFile(d.archivo)
	if os.IsNotExist(err) {
		return []domain.Message{}, nil
	}
	if err != nil {
		return nil, err
	}
	var mensajes []domain.Message
	err = json.Unmarshal(data, &mensajes)
	return mensajes, err
}

func (d *DiskStorage) guardarMensajes(mensajes []domain.Message) error {
	data, err := json.Marshal(mensajes)
	if err != nil {
		return err
	}
	return os.WriteFile(d.archivo, data, 0644)
}

func (d *DiskStorage) Save(message domain.Message) error {
	mensajes, err := d.leerMensajes()
	if err != nil {
		return err
	}
	mensajes = append(mensajes, message)
	return d.guardarMensajes(mensajes)
}

func (d *DiskStorage) FindAll() ([]domain.Message, error) {
	return d.leerMensajes()
}

func (d *DiskStorage) Delete(id string) error {
	mensajes, err := d.leerMensajes()
	if err != nil {
		return err
	}
	for i, msg := range mensajes {
		if msg.ID == id {
			mensajes = append(mensajes[:i], mensajes[i+1:]...)
			return d.guardarMensajes(mensajes)
		}
	}
	return fmt.Errorf("mensaje no encontrado")
}
