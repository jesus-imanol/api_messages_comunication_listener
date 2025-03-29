package repository
import (
	"apimessages/src/messages/domain/entities"

	"github.com/gorilla/websocket"
)

type WebSocketServer interface {
    HandleConnection(conn *websocket.Conn, username string)
    Broadcast(humidity entities.Message)
}
