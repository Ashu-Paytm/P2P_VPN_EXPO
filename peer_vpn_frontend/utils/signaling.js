import { WebSocket } from 'react-native-websocket';
import { SIGNALING_SERVER_URL } from './constants';

let socket;

export const connectToSignalingServer = (onMessage) => {
  socket = new WebSocket(SIGNALING_SERVER_URL);

  socket.onopen = () => console.log('Connected to signaling server');
  socket.onmessage = (message) => onMessage(JSON.parse(message.data));
  socket.onerror = (error) => console.error('WebSocket error:', error);
  socket.onclose = () => console.log('Disconnected from signaling server');
};

export const sendMessage = (message) => {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message));
  } else {
    console.error('WebSocket is not
