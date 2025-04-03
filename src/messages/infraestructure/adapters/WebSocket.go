
package adapters

import (
    "apimessages/src/messages/domain/entities"
    "encoding/json"
    "log"
    "sync"

    "github.com/gorilla/websocket"
)

type WebSocketAdapter struct {
    clients map[string]*websocket.Conn 
    mu      sync.Mutex
}

func NewWebSocketAdapter() *WebSocketAdapter {
    return &WebSocketAdapter{
        clients: make(map[string]*websocket.Conn), 
    }
}

func (ws *WebSocketAdapter) HandleConnection(conn *websocket.Conn, username string) {
    ws.mu.Lock()
    ws.clients[username] = conn 
    ws.mu.Unlock()

    defer func() {
        ws.mu.Lock()
        delete(ws.clients, username) 
        ws.mu.Unlock()
        conn.Close()
    }()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Printf("WebSocket cerrado para usuario %s: %v", username, err)
            break
        }

        log.Printf("Mensaje recibido de usuario %s: %s", username, string(msg))

        if err := conn.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
            log.Printf("Error al enviar pong a usuario %s: %v", username, err)
            break
        }
    }
}

func (ws *WebSocketAdapter) Broadcast(message entities.Message) {
    ws.mu.Lock()
    defer ws.mu.Unlock()

    messageData, err := json.Marshal(message)
    if err != nil {
        log.Println("Error al convertir el mensaje a JSON:", err)
        return
    }

    for username, conn := range ws.clients {
        if username == message.User { 
            err := conn.WriteMessage(websocket.TextMessage, messageData)
            if err != nil {
                log.Printf("Error al enviar mensaje a usuario %s: %v", username, err)
                conn.Close()
                delete(ws.clients, username) 
            }
        }
    }
}