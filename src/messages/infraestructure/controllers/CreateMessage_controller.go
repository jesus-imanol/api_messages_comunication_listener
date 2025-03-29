package controllers

import (
	"apimessages/src/messages/application"
	"apimessages/src/messages/domain/entities"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
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
	var humidity entities.Message
	if err := g.ShouldBindJSON(&humidity); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	messagetoCreate,err := cm.createMessageUsecase.Execute(humidity)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(messagetoCreate)
	g.JSON(http.StatusCreated, messagetoCreate)

}