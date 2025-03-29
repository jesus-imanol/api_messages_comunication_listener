package routers

import (
	middleware "apimessages/src/messages/application/middlewares"
	"apimessages/src/messages/infraestructure/adapters"
	"apimessages/src/messages/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.Engine, webSocketAdapter *adapters.WebSocketAdapter, createmessageController *controllers.CreateMessageController) {
    v1 := r.Group("/v1/message")
    {
        v1.POST("/", createmessageController.CreateMessage)
        v1.GET("/ws", middleware.WebSocketMiddleware(webSocketAdapter))
    }
}