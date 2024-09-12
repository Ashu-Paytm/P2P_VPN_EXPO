package websockets

import "vpn_backend/models"

var vpnSessions = make(map[string]models.VPNSession)

func AddSession(session models.VPNSession) {
	vpnSessions[session.SessionID] = session
}

func GetSession(sessionID string) (models.VPNSession, bool) {
	session, found := vpnSessions[sessionID]
	return session, found
}

func UpdateSession(session models.VPNSession) {
	vpnSessions[session.SessionID] = session
}

func GetAllSessions() []models.VPNSession {
	sessions := make([]models.VPNSession, 0, len(vpnSessions))
	for _, session := range vpnSessions {
		sessions = append(sessions, session)
	}
	return sessions
}
