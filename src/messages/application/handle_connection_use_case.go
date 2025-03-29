package application

import (
	repository "apimessages/src/messages/application/reposository"

	"github.com/gorilla/websocket"
)

type HandleConnectionUseCase struct {
	connection repository.WebSocketServer
}

// Constructor
func NewHandleConnectionUseCase(connectionn repository.WebSocketServer) *HandleConnectionUseCase {
    return &HandleConnectionUseCase{connection:connectionn}
}

// Maneja nuevas conexiones WebSocket
func (uc *HandleConnectionUseCase) Execute(conn *websocket.Conn, username string) {
    uc.connection.HandleConnection(conn, username)
}
