package routers

import (
	"apimessages/src/messages/infraestructure/adapters"
	"apimessages/src/messages/infraestructure/controllers"
	services "apimessages/src/messages/infraestructure/serivces"

	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.Engine, key string, webSocketAdapter *adapters.WebSocketAdapter, createmessageController *controllers.CreateMessageController, initFertilizerController *controllers.InitFertilizerController) {
    v1 := r.Group("/v1/message")
    protectedRouteConsumer := v1.Group("/consumer")
    protectedRouteConsumer.Use(services.RoleMiddleware(key, []string{"controller"}))
    protectedRouteConsumer.POST("/", createmessageController.CreateMessage)
   // protectedRouteTwoUsers := v1.Group("/protected") 
   // protectedRouteTwoUsers.Use(services.RoleMiddleware(key, []string{"normaluser", "premiumuser"}))
    v1.GET("/ws", services.WebSocketMiddleware(webSocketAdapter))
    v1.POST("/messageFertilizer", initFertilizerController.Run)
}