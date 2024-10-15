export const RTCConfiguration = {
    iceServers: [
      {
        urls: 'stun:stun.l.google.com:19302',
      },
      {
        urls: 'turn:your-turn-server-url',
        username: 'your-username',
        credential: 'your-password',
      },
    ],
  };
  
  export const WebRTCOptions = {
    offerToReceiveAudio: true,
    offerToReceiveVideo: true,
  };
  