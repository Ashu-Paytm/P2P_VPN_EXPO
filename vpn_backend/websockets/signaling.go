package websockets

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
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

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		log.Printf("Message from %s: %s", peerID, message)
		s.BroadcastMessage(peerID, message)
	}
}
