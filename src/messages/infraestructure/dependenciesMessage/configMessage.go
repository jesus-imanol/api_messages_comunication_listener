package dependenciesMessage

import (
	"apimessages/src/messages/application"
	"apimessages/src/messages/infraestructure/adapters"
	"apimessages/src/messages/infraestructure/controllers"
	"apimessages/src/messages/infraestructure/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitMessages(r *gin.Engine, webSocketAdapter *adapters.WebSocketAdapter, conecctionMysql *adapters.MySQL) {
	smtpAdapter := adapters.NewSMTP()
    err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading.env file")
    }
	key := os.Getenv("SECRET_KEY")
	createMessageUseCase := application.NewCreateMessageUsecase(conecctionMysql, webSocketAdapter, smtpAdapter)
	createMessageController := controllers.NewCreateMessageController(createMessageUseCase)

	routers.MessageRouter(r, key, webSocketAdapter, createMessageController)
}