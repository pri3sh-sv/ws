package main

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var url = "ws://localhost:8080/ws" // server-url

func connectToWebSocket() {
	for {
		log.Println("Connecting to WebSocket...")

		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Printf("WebSocket connection error: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("WebSocket connected")

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error: %v", err)
				break
			}
			log.Println("WebSocket received:", string(message))
		}

		conn.Close()
		log.Println("Reconnecting in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	connectToWebSocket()
}
