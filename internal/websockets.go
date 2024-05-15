package websockets

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// Check if the received message is binary (assuming blobs are sent as binary data)
		if messageType == websocket.BinaryMessage {
			fmt.Println("Video received")
			fmt.Printf("Data received: %x\n", message)
		}

		// Here you would process the video stream data
		// For simplicity, we're just echoing back the received data
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
