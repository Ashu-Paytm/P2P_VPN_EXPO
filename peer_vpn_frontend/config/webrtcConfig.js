
// STUN and TURN servers for peer-to-peer connection negotiation. STUN (Session Traversal Utilities for NAT) is used to discover the public IP address of a peer, and TURN (Traversal Using Relays around NAT) is used when STUN fails, providing a relay server for communication.

export const RTCConfiguration = {
  iceServers: [
    // STUN server (Google's public STUN server)
    {
      urls: ['stun:bn-turn1.xirsys.com'],
    },
    // TURN server configuration (using your provided static TURN credentials)
    {
      username: 'YI7i6Jwg-FBam7wTUxmbw3z8W6uihOxqtRSOmJYNhhe7r01ZG9aV17h2OAgEAbeoAAAAAGcsx6J0aGVhbG9uZW11c2s=',
      credential: '6f449536-9d10-11ef-a3c6-0242ac140004',
      urls: [
        'turn:bn-turn1.xirsys.com:80?transport=udp',     // UDP transport on port 80
        'turn:bn-turn1.xirsys.com:3478?transport=udp',   // UDP transport on port 3478
        'turn:bn-turn1.xirsys.com:80?transport=tcp',     // TCP transport on port 80
        'turn:bn-turn1.xirsys.com:3478?transport=tcp',   // TCP transport on port 3478
        'turns:bn-turn1.xirsys.com:443?transport=tcp',    // TLS/SSL TCP transport on port 443
        'turns:bn-turn1.xirsys.com:5349?transport=tcp',   // TLS/SSL TCP transport on port 5349
      ],
    },
  ],
};

export const WebRTCOptions = {
  offerToReceiveAudio: true,  // Enable audio reception
  offerToReceiveVideo: true,  // Enable video reception
};

  