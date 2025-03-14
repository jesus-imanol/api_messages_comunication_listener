package dependenciesmessage

import (
	"apimessages/src/humidity/application"
	"apimessages/src/humidity/infraestructure/adapters"
	"apimessages/src/humidity/infraestructure/controllers"
	"apimessages/src/humidity/infraestructure/routers"
	"github.com/gin-gonic/gin"
)

func InitMessages(r *gin.Engine, webSocketAdapter *adapters.WebSocketAdapter , conecctionMysql *adapters.MySQL) {
	//init repository
	createMessageUseCase := application.NewCreateMessageUsecase(conecctionMysql, webSocketAdapter)
	createMessageController := controllers.NewCreateMessageController(createMessageUseCase)

	routers.MessageRouter(r,  webSocketAdapter, createMessageController)
}