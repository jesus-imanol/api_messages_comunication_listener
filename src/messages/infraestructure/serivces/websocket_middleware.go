package services

import (
    "apimessages/src/messages/infraestructure/adapters"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

func WebSocketMiddleware(webSocketAdapter *adapters.WebSocketAdapter) gin.HandlerFunc {
    return func(c *gin.Context) {
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

        username := c.Query("user") 
        if username == "" {
            log.Println("Falta el identificador de usuario")
            conn.Close()
            c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el identificador de usuario"})
            return
        }

        log.Printf("Usuario conectado: %s", username)

        webSocketAdapter.HandleConnection(conn, username)
    }
}