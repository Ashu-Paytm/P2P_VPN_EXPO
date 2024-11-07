// import React, { useState } from 'react';
// import { View, Text, TextInput, Button, StyleSheet } from 'react-native';

// const HostVPN = ({ onCreate }) => {
//   const [sessionID, setSessionID] = useState('');

//   return (
//     <View style={styles.container}>
//       <Text>Host a VPN Session</Text>
//       <TextInput
//         style={styles.input}
//         placeholder="Enter Session ID"
//         value={sessionID}
//         onChangeText={setSessionID}
//       />
//       <Button title="Create VPN" onPress={() => onCreate(sessionID)} />
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { padding: 20 },
//   input: { borderWidth: 1, marginVertical: 10, padding: 8 },
// });

// export default HostVPN;

import React, { useEffect, useState } from 'react';
import { Button, Text, View } from 'react-native';
import useWebSocket from './signaling'; // Import WebSocket hook
import { RTCConfiguration, WebRTCOptions } from '../config/webrtcConfig'; // Import WebRTC config

const HostVPN = () => {
  const { socket, connected } = useWebSocket();
  const [peerConnection, setPeerConnection] = useState(null);

  useEffect(() => {
    if (connected) {
      // Initialize WebRTC PeerConnection when connected to the signaling server
      const pc = new RTCPeerConnection(RTCConfiguration);

      // Handle ICE candidate events
      pc.onicecandidate = (event) => {
        if (event.candidate) {
          console.log('Sending ICE candidate:', event.candidate);
          socket.send(JSON.stringify({ type: 'candidate', candidate: event.candidate }));
        }
      };

      // Set up the local stream (e.g., audio/video)
      // You would typically use `getUserMedia` to get the media stream
      // For example: pc.addStream(localStream);

      setPeerConnection(pc);
    }

    return () => {
      // Clean up WebRTC connection on component unmount
      if (peerConnection) {
        peerConnection.close();
      }
    };
  }, [connected]);

  const createOffer = async () => {
    if (!peerConnection) return;

    try {
      const offer = await peerConnection.createOffer(WebRTCOptions);
      await peerConnection.setLocalDescription(offer);

      console.log('Sending offer:', offer);
      socket.send(JSON.stringify({ type: 'offer', offer }));
    } catch (error) {
      console.error('Error creating offer:', error);
    }
  };

  return (
    <View>
      <Text>Host VPN</Text>
      <Button title="Create Offer" onPress={createOffer} disabled={!connected} />
    </View>
  );
};

export default HostVPN;
