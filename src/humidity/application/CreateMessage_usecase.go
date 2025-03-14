package application

import (
	repositories "apimessages/src/humidity/application/reposositories"
	"apimessages/src/humidity/domain/entities"
	repository "apimessages/src/humidity/domain/repositories"
)
type CreateMessageUsecase struct {
	MessageRepository repository.IMessage
	WebSocketService repositories.WebSocketServer
}

// NewCreateMessageUsecase crea un nuevo caso de uso de creación de mensajes
func NewCreateMessageUsecase(messageRepository repository.IMessage, wsService repositories.WebSocketServer) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageRepository: messageRepository,
		WebSocketService:  wsService,
	}
}

// Execute maneja la creación del mensaje y lo envía a través de WebSocket
func (uc *CreateMessageUsecase) Execute(humidity entities.Humidity) (*entities.Humidity, error) {
	humidityReceive, err := uc.MessageRepository.CreateMessage(humidity)
	if err != nil {
		return nil, err
	}
	uc.WebSocketService.Broadcast(*humidityReceive)

	return humidityReceive, nil
}
