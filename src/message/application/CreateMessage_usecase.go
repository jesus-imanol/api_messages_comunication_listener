package application

import (
	"apimessages/src/message/domain/entities"
	"apimessages/src/message/domain/repositories"
)

type CreateMessageUsecase struct {
	MessageRepository repositories.IMessage
}

func NewCreateMessageUsecase(messageRepository repositories.IMessage) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageRepository: messageRepository,
	}
}

func (uc *CreateMessageUsecase) Execute(typing string, quantity float64, text string ) (*entities.Message,error) {
	message := entities.NewMessage(typing, quantity, text)
	err := uc.MessageRepository.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil

}