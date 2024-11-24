package config

import (
	"encoding/json"
	"log"
	"os"
)

type RTCServer struct {
	URLs       []string `json:"urls"`
	Username   string   `json:"username"`
	Credential string   `json:"credential"`
}

type Config struct {
	ServerPort    string      `json:"server_port"`
	ICEServers    []RTCServer `json:"ice_servers"`
	EncryptionKey string      `json:"encryption_key"`
}

var AppConfig Config

func LoadConfig() {
	// Load from config.json if exists
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&AppConfig); err != nil {
			log.Printf("Error decoding config file: %v", err)
		}
	}

	// Override with environment variables if present
	AppConfig.ServerPort = "8080"
	// AppConfig.ServerPort = getEnv("SERVER_PORT", AppConfig.ServerPort)
	// if AppConfig.ServerPort == "" {
	// 	AppConfig.ServerPort = "8080"
	// }

	// Set default STUN server if none configured
	if len(AppConfig.ICEServers) == 0 {
		AppConfig.ICEServers = []RTCServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		}
	}

	AppConfig.EncryptionKey = "asjdfiosersep"
	// AppConfig.EncryptionKey = getEnv("ENCRYPTION_KEY", AppConfig.EncryptionKey)
	// if AppConfig.EncryptionKey == "" {
	// 	log.Fatal("ENCRYPTION_KEY must be set")
	// }
}
