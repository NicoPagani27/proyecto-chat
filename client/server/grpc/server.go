package grpc

import (
	"fmt"
	"net"

	googlegrpc "google.golang.org/grpc"

	proto "proyecto-chat/client/server/proto"
	"proyecto-chat/domain"
)

type GRPCServer struct {
	send   *domain.SendMessageUseCase
	list   *domain.ListMessagesUseCase
	delete *domain.DeleteMessageUseCase
}

func NewGRPCServer(
	send *domain.SendMessageUseCase,
	list *domain.ListMessagesUseCase,
	delete *domain.DeleteMessageUseCase,
) *GRPCServer {
	return &GRPCServer{send: send, list: list, delete: delete}
}

func (s *GRPCServer) Start(direccion string) error {
	listener, err := net.Listen("tcp", direccion)
	if err != nil {
		return fmt.Errorf("no se pudo escuchar en %s: %w", direccion, err)
	}

	servidorGRPC := googlegrpc.NewServer()
	proto.RegisterChatServiceServer(servidorGRPC, &manejadores{
		enviar:   s.send,
		listar:   s.list,
		eliminar: s.delete,
	})

	return servidorGRPC.Serve(listener)
}
