package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "proyecto-chat/client/server/proto"
	"proyecto-chat/domain"
)

type manejadores struct {
	proto.UnimplementedChatServiceServer
	enviar   *domain.SendMessageUseCase
	listar   *domain.ListMessagesUseCase
	eliminar *domain.DeleteMessageUseCase
}

func (m *manejadores) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	mensaje, err := m.enviar.Execute(req.Author, req.Text)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &proto.SendMessageResponse{
		Message: domainAProto(mensaje),
	}, nil
}

func (m *manejadores) ListMessages(ctx context.Context, req *proto.ListMessagesRequest) (*proto.ListMessagesResponse, error) {
	mensajes, err := m.listar.Execute()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var protoMensajes []*proto.ChatMessage
	for _, msg := range mensajes {
		protoMensajes = append(protoMensajes, domainAProto(msg))
	}
	return &proto.ListMessagesResponse{Messages: protoMensajes}, nil
}

func (m *manejadores) DeleteMessage(ctx context.Context, req *proto.DeleteMessageRequest) (*proto.DeleteMessageResponse, error) {
	err := m.eliminar.Execute(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &proto.DeleteMessageResponse{}, nil
}

func domainAProto(msg domain.Message) *proto.ChatMessage {
	return &proto.ChatMessage{
		Id:        msg.ID,
		Author:    msg.Author,
		Text:      msg.Text,
		Timestamp: msg.Timestamp.Format(time.RFC3339),
	}
}
