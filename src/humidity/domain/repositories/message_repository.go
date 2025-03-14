package repositories

import "apimessages/src/humidity/domain/entities"

type IMessage interface {
	CreateMessage(humidity entities.Humidity) (*entities.Humidity, error)
}