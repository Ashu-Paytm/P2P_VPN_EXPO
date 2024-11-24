package models

import (
	"encoding/json"
	"time"
)

type Peer struct {
	PeerID   string    `json:"peer_id"`
	IsActive bool      `json:"is_active"`
	JoinedAt time.Time `json:"joined_at"`
	LastSeen time.Time `json:"last_seen"`
}

type VPNSession struct {
	SessionID   string    `json:"session_id"`
	CreatedAt   time.Time `json:"created_at"`
	Peers       []Peer    `json:"peers"`
	IsEncrypted bool      `json:"is_encrypted"`
}

type SignalingMessage struct {
	Type      string          `json:"type"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Payload   json.RawMessage `json:"payload"`
	SessionID string          `json:"session_id"`
}
