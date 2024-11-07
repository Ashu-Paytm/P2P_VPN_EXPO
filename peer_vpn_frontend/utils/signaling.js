import { useEffect, useState } from 'react';
import { SIGNALING_SERVER_URL } from './constants';

const useWebSocket = () => {
  const [socket, setSocket] = useState(null);
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    const ws = new WebSocket(SIGNALING_SERVER_URL);

    // Set up WebSocket event listeners
    ws.onopen = () => {
      console.log('Connected to signaling server');
      setConnected(true);
    };

    ws.onclose = () => {
      console.log('Disconnected from signaling server');
      setConnected(false);
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log('Received message:', data);

      // Handle signaling messages (e.g., offer, answer, ice candidate)
      if (data.type === 'config') {
        // Handle the config (TURN/STUN servers)
        console.log('Received TURN/STUN config:', data.iceServers);
      } else {
        // Handle other signaling messages (offer, answer, ICE candidates)
        // For example: data.type === 'offer' or 'answer'
        console.log('Signaling message:', data);
      }
    };

    // Clean up WebSocket on component unmount
    return () => {
      ws.close();
    };
  }, []);

  return { socket, connected };
};

export default useWebSocket;

      
