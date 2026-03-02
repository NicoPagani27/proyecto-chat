package main

import (
	"flag"
	"log"

	"proyecto-chat/client/server/grpc"
	"proyecto-chat/client/server/rest"
	"proyecto-chat/client/storage"
	"proyecto-chat/domain"
)

func main() {
	// Flags
	storageFlag := flag.String("storage", "memory", "memory | disk | sqlite")
	serverFlag := flag.String("server", "rest", "rest | grpc | dual")
	restPort := flag.String("rest-port", ":8080", "puerto REST")
	grpcPort := flag.String("grpc-port", ":9090", "puerto gRPC")
	flag.Parse()

	// Crear el repositorio según el flag
	var repo domain.MessageRepository
	switch *storageFlag {
	case "memory":
		repo = storage.NewMemoryStorage()
	case "disk":
		repo = storage.NewDiskStorage("data/messages.json")
	case "sqlite":
		s, err := storage.NewSQLiteStorage("data/messages.db")
		if err != nil {
			log.Fatalf("error al abrir SQLite: %v", err)
		}
		repo = s
	default:
		log.Fatalf("storage inválido: %s", *storageFlag)
	}

	// Crear los use cases
	send := domain.NewSendMessageUseCase(repo)
	list := domain.NewListMessagesUseCase(repo)
	delete := domain.NewDeleteMessageUseCase(repo)

	// Iniciar el servidor según el flag
	switch *serverFlag {
	case "rest":
		log.Printf("Iniciando servidor REST en %s", *restPort)
		restServer := rest.NewRESTServer(send, list, delete)
		if err := restServer.Start(*restPort); err != nil {
			log.Fatalf("error en servidor REST: %v", err)
		}
	case "grpc":
		log.Printf("Iniciando servidor gRPC en %s", *grpcPort)
		grpcServer := grpc.NewGRPCServer(send, list, delete)
		if err := grpcServer.Start(*grpcPort); err != nil {
			log.Fatalf("error en servidor gRPC: %v", err)
		}
	case "dual":
		log.Printf("Iniciando servidor REST en %s y gRPC en %s", *restPort, *grpcPort)
		restServer := rest.NewRESTServer(send, list, delete)
		grpcServer := grpc.NewGRPCServer(send, list, delete)
		go func() {
			if err := grpcServer.Start(*grpcPort); err != nil {
				log.Fatalf("error en servidor gRPC: %v", err)
			}
		}()
		if err := restServer.Start(*restPort); err != nil {
			log.Fatalf("error en servidor REST: %v", err)
		}
	default:
		log.Fatalf("server inválido: %s", *serverFlag)
	}
}
