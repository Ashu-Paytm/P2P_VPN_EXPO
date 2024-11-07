package websockets

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"vpn-backend/config" // Import your config package to access RTCConfiguration
)

// Peer stores the connection and ID of each peer
type Peer struct {
	Conn   *websocket.Conn
	PeerID string
}

// SignalingServer manages all peers and connections
type SignalingServer struct {
	Peers map[string]*Peer
	lock  sync.Mutex
}

// NewSignalingServer creates a new signaling server
func NewSignalingServer() *SignalingServer {
	return &SignalingServer{
		Peers: make(map[string]*Peer),
	}
}

// AddPeer adds a peer to the signaling server
func (s *SignalingServer) AddPeer(peerID string, conn *websocket.Conn) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Peers[peerID] = &Peer{
		Conn:   conn,
		PeerID: peerID,
	}

	log.Printf("Peer added: %s", peerID)
}

// RemovePeer removes a peer from the signaling server
func (s *SignalingServer) RemovePeer(peerID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if peer, ok := s.Peers[peerID]; ok {
		peer.Conn.Close()
		delete(s.Peers, peerID)
		log.Printf("Peer removed: %s", peerID)
	}
}

// BroadcastMessage broadcasts a message to all peers except the sender
func (s *SignalingServer) BroadcastMessage(senderID string, message []byte) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for peerID, peer := range s.Peers {
		if peerID != senderID {
			if err := peer.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Error sending message to peer %s: %v", peerID, err)
			}
		}
	}
}

// HandleSignaling handles signaling messages from peers
func (s *SignalingServer) HandleSignaling(peerID string, conn *websocket.Conn) {
	defer s.RemovePeer(peerID)

	s.AddPeer(peerID, conn)

	// Prepare the message that will include TURN/STUN info
	turnStunConfig := config.RTCConfiguration.ICEServers

	// Send the TURN/STUN configuration to the peer when they first connect
	initialMessage := map[string]interface{}{
		"type":       "config",
		"iceServers": turnStunConfig,
	}
	if err := conn.WriteJSON(initialMessage); err != nil {
		log.Printf("Error sending TURN/STUN config to %s: %v", peerID, err)
	}

	// Handle further signaling messages (offer, answer, ICE candidates)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		log.Printf("Message from %s: %s", peerID, message)

		// Include TURN/STUN config in the signaling message
		messageWithConfig := map[string]interface{}{
			"message":    string(message),
			"iceServers": turnStunConfig,
		}

		// Broadcast the message to all other peers
		s.BroadcastMessage(peerID, messageWithConfig)
	}
}
