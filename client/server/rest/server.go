package rest

import (
	"net/http"

	"proyecto-chat/domain"
)

type RESTServer struct {
	send   *domain.SendMessageUseCase
	list   *domain.ListMessagesUseCase
	delete *domain.DeleteMessageUseCase
}

func NewRESTServer(
	send *domain.SendMessageUseCase,
	list *domain.ListMessagesUseCase,
	delete *domain.DeleteMessageUseCase,
) *RESTServer {
	return &RESTServer{send: send, list: list, delete: delete}
}

func (s *RESTServer) Start(direccion string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/messages", s.handleMessages)
	mux.HandleFunc("/messages/", s.handleMessageByID)

	fs := http.FileServer(http.Dir("frontend-quantex-chat"))
	mux.Handle("/", fs)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		mux.ServeHTTP(w, r)
	})

	return http.ListenAndServe(direccion, handler)
}
