package services

import (
	"errors"
	"vpn_backend/models"
)

var vpnSessions = make(map[string]models.VPNSession) // Simple in-memory storage

func CreateVPN(vpn models.VPNSession) (string, error) {
	if _, exists := vpnSessions[vpn.SessionID]; exists {
		return "", errors.New("session already exists")
	}
	vpnSessions[vpn.SessionID] = vpn
	return vpn.SessionID, nil
}

func JoinVPN(sessionID, peerID string) error {
	session, exists := vpnSessions[sessionID]
	if !exists {
		return errors.New("session not found")
	}

	session.Peers = append(session.Peers, models.Peer{PeerID: peerID, IsActive: true})
	vpnSessions[sessionID] = session
	return nil
}
