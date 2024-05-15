package websockets

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Connection struct {
	Conn *websocket.Conn
}

var Connections sync.Map = sync.Map{}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	// Generate a unique ID for this connection
	id := generateID()

	// Store the connection in the global map
	Connections.Store(id, &Connection{Conn: conn})

	// Broadcast the connection ID to all clients so they know who is connected
	broadcastMessage(id)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// Process the message and broadcast it to all clients
		processAndBroadcastMessage(id, message)
	}
}

func generateID() string {
	// Simple ID generation for demonstration purposes
	return "client-" + strconv.FormatInt(time.Now().UnixNano(), 10)
}

func broadcastMessage(id string) {
	// Iterate over all connections and broadcast the message
	Connections.Range(func(key, value interface{}) bool {
		c := value.(*Connection)
		err := c.Conn.WriteMessage(websocket.TextMessage, []byte("New client connected: "+id))
		if err != nil {
			log.Println("broadcast:", err)
		}
		return true
	})
}

func processAndBroadcastMessage(clientID string, message []byte) {
	// Process the message here
	// For demonstration, we're just broadcasting the message back to all clients

	Connections.Range(func(key, value interface{}) bool {
		c := value.(*Connection)
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("broadcast:", err)
		}
		return true
	})
}
