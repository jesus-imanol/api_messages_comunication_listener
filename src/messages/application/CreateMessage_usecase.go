package application

import (
	repository "apimessages/src/messages/application/reposository"
	"apimessages/src/messages/domain/entities"
	"apimessages/src/messages/domain/repositories"
)

type CreateMessageUsecase struct {
	MessageRepository repositories.IMessage
	WebSocketService  repository.WebSocketServer
	SMTPRepository    repository.ISmtp
}

func NewCreateMessageUsecase(messageRepository repositories.IMessage, wsService repository.WebSocketServer, smtpRepository repository.ISmtp) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageRepository: messageRepository,
		WebSocketService:  wsService,
		SMTPRepository:    smtpRepository,
	}
}

func (uc *CreateMessageUsecase) Execute(message entities.Message) (*entities.Message, error) {
	humidityReceive, err := uc.MessageRepository.CreateMessage(message)
	if err != nil {
		return nil, err
	}

	uc.WebSocketService.Broadcast(*humidityReceive)

	gmail, err := uc.MessageRepository.GetGmailByUserName(message.User)
	if err != nil {
		return nil, err
	}

	if message.Type == "humidity" {
		if message.Text == "¡ALERTA! Suelo seco"{
			err := uc.SMTPRepository.CaseError(message, gmail)
			if err != nil {
				return nil, err
			}
		}
	}

	if message.Type == "temperature" {
		if message.Text == "¡ALERTA! Temperatura baja" || message.Text == "¡ALERTA! Temperatura alta" {
			err := uc.SMTPRepository.CaseError(message, gmail)
			if err != nil {
				return nil, err
			}
		}
	}

	return humidityReceive, nil
}