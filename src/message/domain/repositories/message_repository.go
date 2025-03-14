package repositories

import "apimessages/src/message/domain/entities"

type IMessage interface {
	CreateMessage(message *entities.Message) error
}