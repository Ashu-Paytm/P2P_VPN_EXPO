import { RTCPeerConnection } from 'react-native-webrtc';
import { RTCConfiguration } from '../config/webrtcConfig';

export const createPeerConnection = (onIceCandidate, onTrack) => {
  const peerConnection = new RTCPeerConnection(RTCConfiguration);

  peerConnection.onicecandidate = (event) => {
    if (event.candidate) {
      onIceCandidate(event.candidate);
    }
  };

  peerConnection.ontrack = (event) => {
    onTrack(event.streams[0]);
  };

  return peerConnection;
};

export const createOffer = async (peerConnection) => {
  const offer = await peerConnection.createOffer();
  await peerConnection.setLocalDescription(offer);
  return offer;
};

export const createAnswer = async (peerConnection) => {
  const answer = await peerConnection.createAnswer();
  await peerConnection.setLocalDescription(answer);
  return answer;
};

export const addIceCandidate = async (peerConnection, candidate) => {
  await peerConnection.addIceCandidate(candidate);
};
