package application

import (
	repository "apimessages/src/messages/application/reposository"
	"apimessages/src/messages/domain/entities"
)

type BroadcastMessageUseCase struct {
	server repository.WebSocketServer
}

// Constructor
func NewBroadcastMessageUseCase(server repository.WebSocketServer) *BroadcastMessageUseCase {
	return &BroadcastMessageUseCase{server: server}
}

// Envía mensajes a todos los clientes conectados
func (uc *BroadcastMessageUseCase) Execute(message entities.Message) {
	uc.server.Broadcast(message)
}
