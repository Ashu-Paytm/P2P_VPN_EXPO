package services

import (
	"crypto/rand"
	"encoding/hex"
	"time"
	"vpn_backend/config"
	"vpn_backend/models"
)

type VPNService struct {
	signalingServer *websockets.SignalingServer
}

func NewVPNService(server *websockets.SignalingServer) *VPNService {
	return &VPNService{
		signalingServer: server,
	}
}

func (s *VPNService) CreateSession() (*models.VPNSession, error) {
	sessionID := generateSessionID()
	session := &models.VPNSession{
		SessionID:   sessionID,
		CreatedAt:   time.Now(),
		Peers:       make([]models.Peer, 0),
		IsEncrypted: true,
	}

	return session, nil
}

func generateSessionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
