package dependenciesMessage

import (
	"apimessages/src/messages/application"
	"apimessages/src/messages/infraestructure/adapters"
	"apimessages/src/messages/infraestructure/controllers"
	"apimessages/src/messages/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func InitMessages(r *gin.Engine, webSocketAdapter *adapters.WebSocketAdapter , conecctionMysql *adapters.MySQL) {
	//init repository
	createMessageUseCase := application.NewCreateMessageUsecase(conecctionMysql, webSocketAdapter)
	createMessageController := controllers.NewCreateMessageController(createMessageUseCase)

	routers.MessageRouter(r,  webSocketAdapter, createMessageController)
}