package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *RESTServer) handleMessages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handleSend(w, r)
	case http.MethodGet:
		s.handleList(w, r)
	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

func (s *RESTServer) handleMessageByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/messages/")
	if id == "" {
		http.Error(w, "id requerido", http.StatusBadRequest)
		return
	}
	err := s.delete.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *RESTServer) handleSend(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Author string `json:"author"`
		Text   string `json:"text"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "request inválida", http.StatusBadRequest)
		return
	}
	mensaje, err := s.send.Execute(body.Author, body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mensaje)
}

func (s *RESTServer) handleList(w http.ResponseWriter, r *http.Request) {
	mensajes, err := s.list.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mensajes)
}
