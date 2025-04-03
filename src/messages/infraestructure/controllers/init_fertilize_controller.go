package controllers

import (
	"apimessages/src/messages/application"
	"apimessages/src/messages/domain/entities"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InitFertilizerController struct {
	useCase *application.InitFertilizerUseCase
}

func NewInitFertilizerController(initFertilizerUseCase *application.InitFertilizerUseCase) *InitFertilizerController {
	return &InitFertilizerController{useCase: initFertilizerUseCase}
}

func (ifc *InitFertilizerController) Run(g *gin.Context) {
    var messageFertilizer entities.MessageFertilizer
    if err := g.ShouldBindJSON(&messageFertilizer); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	messageToPublish, err := ifc.useCase.Execute(ctx, messageFertilizer)
	if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	g.JSON(http.StatusOK, gin.H{"data": messageToPublish})
}