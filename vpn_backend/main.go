package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"vpn_backend/config"
	"vpn_backend/services"
	"vpn_backend/websockets"
)

func main() {
	// Initialize configuration
	config.LoadConfig()

	// Initialize signaling server
	websockets.InitSignalingServer()

	// Initialize VPN service
	vpnService := services.NewVPNService(websockets.Server)

	// Setup routes
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/create-session", handleCreateSession)

	// Start server
	serverAddr := ":" + config.AppConfig.ServerPort
	log.Printf("Starting server on %s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	peerID := r.URL.Query().Get("peerID")
	if peerID == "" {
		conn.Close()
		return
	}

	websockets.Server.HandleConnection(conn, peerID)
}
