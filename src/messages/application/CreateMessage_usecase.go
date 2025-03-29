package application

import (
	repository "apimessages/src/messages/application/reposository"
	"apimessages/src/messages/domain/entities"
	"apimessages/src/messages/domain/repositories"
)
type CreateMessageUsecase struct {
	MessageRepository repositories.IMessage
	WebSocketService repository.WebSocketServer
}

// NewCreateMessageUsecase crea un nuevo caso de uso de creación de mensajes
func NewCreateMessageUsecase(messageRepository repositories.IMessage, wsService repository.WebSocketServer) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageRepository: messageRepository,
		WebSocketService:  wsService,
	}
}

// Execute maneja la creación del mensaje y lo envía a través de WebSocket
func (uc *CreateMessageUsecase) Execute(message entities.Message) (*entities.Message, error) {
	humidityReceive, err := uc.MessageRepository.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	uc.WebSocketService.Broadcast(*humidityReceive)

	return humidityReceive, nil
}
