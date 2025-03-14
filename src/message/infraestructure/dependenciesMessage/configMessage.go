package dependenciesmessage

import (
	"apimessages/src/message/application"
	"apimessages/src/message/infraestructure/adapters"
	"apimessages/src/message/infraestructure/controllers"
	"apimessages/src/message/infraestructure/routers"

	"github.com/gin-gonic/gin"
)

func InitMessages(r *gin.Engine, conecctionMysql *adapters.MySQL) {
	//init repository
	createMessageUseCase := application.NewCreateMessageUsecase(conecctionMysql)
	createMessageController := controllers.NewCreateMessageController(createMessageUseCase)

	routers.MessageRouter(r, createMessageController)
}