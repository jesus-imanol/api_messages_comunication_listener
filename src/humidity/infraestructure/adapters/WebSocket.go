package adapters

import (
	"apimessages/src/humidity/domain/entities"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketAdapter struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewWebSocketAdapter() *WebSocketAdapter {
	return &WebSocketAdapter{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (ws *WebSocketAdapter) HandleConnection(conn *websocket.Conn) {
	ws.mu.Lock()
	ws.clients[conn] = true
	ws.mu.Unlock()

	defer func() {
		ws.mu.Lock()
		delete(ws.clients, conn)
		ws.mu.Unlock()
		conn.Close()
	}()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket cerrado:", err)
			break
		}

		log.Printf("Mensaje recibido: %s", msg)

		// Enviar un mensaje de "pong" para mantener la conexión activa
		if err := conn.WriteMessage(messageType, []byte("pong")); err != nil {
			log.Println("Error al enviar pong:", err)
			break
		}
	}
}


func (ws *WebSocketAdapter) Broadcast(humidity entities.Humidity) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	messageData, err := json.Marshal(humidity)
	if err != nil {
		log.Println("Error al convertir el mensaje a JSON:", err)
		return
	}	

	for conn := range ws.clients {
		err := conn.WriteMessage(websocket.TextMessage, messageData)
		if err != nil {
			log.Println("Error al enviar mensaje:", err)
			conn.Close()
			delete(ws.clients, conn)
		}
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, wsAdapter *WebSocketAdapter) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}


	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al establecer conexión:", err)
		return
	}
	
	wsAdapter.HandleConnection(conn)
}
