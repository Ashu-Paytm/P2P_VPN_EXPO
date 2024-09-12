package services

import (
	"github.com/gorilla/websocket"
	"log"
)

func HandleSignaling(conn *websocket.Conn, peerID, sessionID string) error {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read message error: %v", err)
			return err
		}

		// Here you would relay messages between peers
		log.Printf("Received signaling message from %s in session %s: %s", peerID, sessionID, string(message))

		// For simplicity, just echo back the message for now
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Write message error: %v", err)
			return err
		}
	}
}
