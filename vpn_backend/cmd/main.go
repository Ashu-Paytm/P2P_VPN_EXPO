package main

import (
	"log"
	"net/http"
	"vpn_backend/api"
	"vpn_backend/config"
)

func main() {
	config.LoadConfig() // Load configuration
	http.HandleFunc("/create-vpn", api.CreateVPNHandler)
	http.HandleFunc("/join-vpn", api.JoinVPNHandler)
	http.HandleFunc("/ws", api.SignalingHandler) // WebSocket handler for signaling

	log.Printf("Server started at port %s", config.AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, nil))
}
