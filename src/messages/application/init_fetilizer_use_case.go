package application

import (
	"apimessages/src/messages/domain/entities"
	"apimessages/src/messages/domain/repositories"
	"context"
)

type InitFertilizerUseCase struct {
	repository repositories.IMessageFertilizerRepository
}

func NewInitFertilizer(repository repositories.IMessageFertilizerRepository) *InitFertilizerUseCase {
    return &InitFertilizerUseCase{repository}
}

func (i *InitFertilizerUseCase) Execute(ctx context.Context,messageFertilizer entities.MessageFertilizer) (*entities.MessageFertilizer, error) {
	messageFertilizerReceive, err := i.repository.InitFertilizer(ctx, messageFertilizer)
    if err != nil {
        return nil, err
    }
    return messageFertilizerReceive, nil
}