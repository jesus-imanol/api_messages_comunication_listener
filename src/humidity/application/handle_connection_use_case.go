package application

import (
	repositories "apimessages/src/humidity/application/reposositories"

	"github.com/gorilla/websocket"
)

type HandleConnectionUseCase struct {
	connection repositories.WebSocketServer
}

// Constructor
func NewHandleConnectionUseCase(connectionn repositories.WebSocketServer) *HandleConnectionUseCase {
    return &HandleConnectionUseCase{connection:connectionn }
}

// Maneja nuevas conexiones WebSocket
func (uc *HandleConnectionUseCase) Execute(conn *websocket.Conn) {
    uc.connection.HandleConnection(conn)
}
