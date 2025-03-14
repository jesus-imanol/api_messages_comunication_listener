package controllers

import (
	"apimessages/src/message/application"
	"apimessages/src/message/domain/entities"

	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateMessageController struct {
	createMessageUsecase *application.CreateMessageUsecase
}
func NewCreateMessageController(createMessageUsecase *application.CreateMessageUsecase) *CreateMessageController {
	return &CreateMessageController{
		createMessageUsecase: createMessageUsecase,
	}
}
func (cm *CreateMessageController) CreateMessage(g *gin.Context) {
	var message entities.Message
	if err := g.ShouldBindJSON(&message); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	messagetoCreate,err := cm.createMessageUsecase.Execute(message.Type, message.Quantity, message.Text)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, messagetoCreate)

}