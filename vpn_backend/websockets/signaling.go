package websockets

import (
	// "encoding/json"
	// "os"
	// "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
	"vpn_backed/services/vpnService"
	"vpn_backend/config"
	"vpn_backend/models"
)

type SignalingServer struct {
	sessions map[string]*VPNSession
	peers    map[string]*websocket.Conn
	mutex    sync.RWMutex
}

var Server *SignalingServer

func InitSignalingServer() {
	Server = &SignalingServer{
		sessions: make(map[string]*VPNSession),
		peers:    make(map[string]*websocket.Conn),
	}
}

func (s *SignalingServer) HandleConnection(conn *websocket.Conn, peerID string) {
	s.mutex.Lock()
	s.peers[peerID] = conn
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.peers, peerID)
		s.mutex.Unlock()
		conn.Close()
	}()

	for {
		var msg SignalingMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		switch msg.Type {
		case "offer", "answer", "ice-candidate":
			s.relayMessage(msg)
		case "join":
			s.handleJoin(msg, peerID)
		}
	}
}

func (s *SignalingServer) relayMessage(msg SignalingMessage) {
	s.mutex.RLock()
	targetConn, exists := s.peers[msg.To]
	s.mutex.RUnlock()

	if exists {
		err := targetConn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error sending message to peer %s: %v", msg.To, err)
		}
	}
}
