package routers

import (
	"apimessages/src/humidity/infraestructure/adapters"
	"apimessages/src/humidity/infraestructure/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func MessageRouter(r *gin.Engine,webSocketAdapter *adapters.WebSocketAdapter , createmessageController *controllers.CreateMessageController) {
	v1 := r.Group("/v1/message")
	{
		v1.POST("/", createmessageController.CreateMessage)
		r.GET("/ws", func(c *gin.Context) {
			upgrader := websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true 
				},
			}
	
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println("Error al realizar upgrade:", err)
				return
			}

			webSocketAdapter.HandleConnection(conn)
		})
	}
}
