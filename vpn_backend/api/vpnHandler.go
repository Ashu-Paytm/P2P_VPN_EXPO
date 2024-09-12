package api

import (
	"encoding/json"
	"net/http"
	"vpn_backend/models"
	"vpn_backend/services"
)

func CreateVPNHandler(w http.ResponseWriter, r *http.Request) {
	var vpn models.VPNSession
	if err := json.NewDecoder(r.Body).Decode(&vpn); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	sessionID, err := services.CreateVPN(vpn)
	if err != nil {
		http.Error(w, "Unable to create VPN", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"sessionID": sessionID})
}

func JoinVPNHandler(w http.ResponseWriter, r *http.Request) {
	var request models.JoinVPNRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := services.JoinVPN(request.SessionID, request.PeerID)
	if err != nil {
		http.Error(w, "Unable to join VPN", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "joined"})
}
