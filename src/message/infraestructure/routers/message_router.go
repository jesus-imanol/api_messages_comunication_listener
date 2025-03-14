package routers

import (
	"apimessages/src/message/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.Engine, createmessageController *controllers.CreateMessageController) {
	v1 := r.Group("/v1/message")
	{
		v1.POST("/", createmessageController.CreateMessage)
	}
}
