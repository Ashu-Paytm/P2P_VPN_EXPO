package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"vpn_backend/models"
	"vpn_backend/websockets"
)

// Create a new VPN session
func CreateVPNSession(w http.ResponseWriter, r *http.Request) {
	var session models.VPNSession
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize an empty list of peers
	session.Peers = make([]models.Peer, 0)

	// Store session in memory (you can later store this in a database)
	websockets.AddSession(session)

	log.Printf("VPN session created with ID: %s", session.SessionID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(session)
}

// Join an existing VPN session
func JoinVPNSession(w http.ResponseWriter, r *http.Request) {
	var joinRequest models.JoinVPNRequest
	err := json.NewDecoder(r.Body).Decode(&joinRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve session
	session, found := websockets.GetSession(joinRequest.SessionID)
	if !found {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	// Add peer to the session
	peer := models.Peer{
		PeerID:   joinRequest.PeerID,
		IsActive: true,
	}
	session.Peers = append(session.Peers, peer)

	// Update session
	websockets.UpdateSession(session)

	log.Printf("Peer %s joined session %s", joinRequest.PeerID, joinRequest.SessionID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)
}

// List all active VPN sessions
func ListSessions(w http.ResponseWriter, r *http.Request) {
	sessions := websockets.GetAllSessions()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessions)
}
