package repositories

import "apimessages/src/messages/domain/entities"

type IMessage interface {
	CreateMessage(message entities.Message) (*entities.Message, error);
	GetGmailByUserName(username string) (string, error);
}