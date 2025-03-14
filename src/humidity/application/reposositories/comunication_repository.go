package repositories

import (
	"apimessages/src/humidity/domain/entities"

	"github.com/gorilla/websocket"
)

type WebSocketServer interface {
    HandleConnection(conn *websocket.Conn)
    Broadcast(humidity entities.Humidity)
}
