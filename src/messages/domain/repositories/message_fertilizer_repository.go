package repositories

import (
	"apimessages/src/messages/domain/entities"

	"golang.org/x/net/context"
)

type IMessageFertilizerRepository interface {
	InitFertilizer(ctx context.Context, messageFertilizer entities.MessageFertilizer) (*entities.MessageFertilizer, error)
}