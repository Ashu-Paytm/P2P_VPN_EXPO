package models

// JoinVPNRequest represents a request to join a VPN session
type JoinVPNRequest struct {
	SessionID string
	PeerID    string
}

// Peer represents a single peer in the VPN session
type Peer struct {
	PeerID   string
	IsActive bool
}

// VPNSession represents a peer-hosted VPN session
type VPNSession struct {
	SessionID string
	Peers     []Peer
}
