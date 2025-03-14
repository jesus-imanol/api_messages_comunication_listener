package application

import (
	repositories "apimessages/src/humidity/application/reposositories"
	"apimessages/src/humidity/domain/entities"
)

type BroadcastMessageUseCase struct {
	server repositories.WebSocketServer
}

// Constructor
func NewBroadcastMessageUseCase(server repositories.WebSocketServer) *BroadcastMessageUseCase {
	return &BroadcastMessageUseCase{server: server}
}

// Envía mensajes a todos los clientes conectados
func (uc *BroadcastMessageUseCase) Execute(message entities.Humidity) {
	uc.server.Broadcast(message)
}
