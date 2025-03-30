package application

import (
	"sync"
	"time"

	repository "apimessages/src/messages/application/reposository"
	"apimessages/src/messages/domain/entities"
	"apimessages/src/messages/domain/repositories"

)

var erroresEnviados = make(map[entities.Message]time.Time)
var mutex sync.Mutex 

type CreateMessageUsecase struct {
	MessageRepository repositories.IMessage
	WebSocketService  repository.WebSocketServer
	SMTPRepository repository.ISmtp
}

func NewCreateMessageUsecase(messageRepository repositories.IMessage, wsService repository.WebSocketServer) *CreateMessageUsecase {
	return &CreateMessageUsecase{
		MessageRepository: messageRepository,
		WebSocketService:  wsService,
	}
}

func (uc *CreateMessageUsecase) Execute(message entities.Message) (*entities.Message, error) {
	humidityReceive, err := uc.MessageRepository.CreateMessage(message)
	if err != nil {
		return nil, err
	}

	uc.WebSocketService.Broadcast(*humidityReceive)
   gmail, err:= uc.MessageRepository.GetGmailByUserName(message.User)
    if err != nil {
		return nil, err
	}
	if message.Type == "humidity" {
		if(message.Text == ""){
			err := uc.SMTPRepository.CaseError(message, gmail)
			if err != nil {
				return nil, err
			}
		}
	}
    if(message.Type == "temperature"){
		if(message.Quantity < 10 || message.Quantity > 30){
			err := uc.SMTPRepository.CaseError(message, gmail)
            if err != nil {
                return nil, err
            }
        }
	}
	return humidityReceive, nil
}
