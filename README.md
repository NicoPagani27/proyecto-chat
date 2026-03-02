# proyecto-chat

Aplicación de chat desarrollada con Go (backend) y Svelte (frontend). Implementa clean architecture y permite elegir entre diferentes tipos de almacenamiento y protocolos de comunicación.

## Características

- Clean architecture con separación de capas
- Soporta REST y gRPC
- Almacenamiento configurable: memoria, disco (JSON) o SQLite  
- Frontend con SvelteKit
- Tests en la capa de dominio

## Arquitectura

El proyecto está estructurado en capas siguiendo los principios de clean architecture. Las dependencias apuntan hacia adentro, donde el dominio no conoce detalles de implementación.

**Casos de uso implementados:**
- Enviar mensaje
- Listar mensajes
- Eliminar mensaje

## Instalación

Necesitas Go 1.25.6+ y Node.js 18+ para el frontend.

```bash
# Backend
go mod download

# Frontend
cd frontend-quantex-chat
npm install
```

Si modificas el proto, regenera el código con:
```bash
protoc --go_out=. --go-grpc_out=. client/server/proto/chat.proto
```

## Uso

El servidor acepta estos flags:

- `-storage`: `memory`, `disk` o `sqlite` (default: `memory`)
- `-server`: `rest`, `grpc` o `dual` (default: `rest`)
- `-rest-port`: puerto REST (default: `:8080`)
- `-grpc-port`: puerto gRPC (default: `:9090`)

### Ejemplos

```bash
# REST + memoria
go run client/main.go -storage=memory -server=rest

# gRPC + SQLite
go run client/main.go -storage=sqlite -server=grpc -grpc-port=:9090

# Dual (REST + gRPC) + disco
go run client/main.go -storage=disk -server=dual
```

## API REST

**POST /messages** - Enviar mensaje
```json
{
  "author": "nombre",
  "text": "contenido"
}
```

**GET /messages** - Listar todos los mensajes

**DELETE /messages/{id}** - Eliminar mensaje

## API gRPC

Ver `chat.proto` para los métodos disponibles (SendMessage, ListMessages, DeleteMessage).

## Estructura

```
proyecto-chat/
├── client/
│   ├── main.go
│   ├── server/
│   │   ├── grpc/
│   │   ├── proto/
│   │   └── rest/
│   └── storage/
├── domain/
│   ├── message.go
│   ├── repository.go
│   └── usecases.go
├── frontend-quantex-chat/
└── data/
```

## Stack

- Go 1.25.6
- gRPC + Protocol Buffers
- SQLite (modernc.org/sqlite)
- Svelte 5 + SvelteKit

## Testing

```bash
go test ./domain/... -v
go test ./... -v
```

## Notas

- storage=memory es volátil
- storage=disk guarda en `data/messages.json`
- storage=sqlite usa `data/messages.db`
- Crea el directorio `data/` antes de usar disk o sqlite
- El modo dual corre REST y gRPC en paralelo