package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"vpn_backend/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func SignalingHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade websocket: %v", err)
		return
	}
	defer conn.Close()

	peerID := r.URL.Query().Get("peerID")
	sessionID := r.URL.Query().Get("sessionID")
	if err := services.HandleSignaling(conn, peerID, sessionID); err != nil {
		log.Printf("Signaling error: %v", err)
	}
}
