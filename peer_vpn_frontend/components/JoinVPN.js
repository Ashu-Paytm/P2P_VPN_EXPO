// import React, { useState } from 'react';
// import { View, Text, TextInput, Button, StyleSheet } from 'react-native';

// const JoinVPN = ({ onJoin }) => {
//   const [sessionID, setSessionID] = useState('');

//   return (
//     <View style={styles.container}>
//       <Text>Join a VPN Session</Text>
//       <TextInput
//         style={styles.input}
//         placeholder="Enter Session ID"
//         value={sessionID}
//         onChangeText={setSessionID}
//       />
//       <Button title="Join VPN" onPress={() => onJoin(sessionID)} />
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { padding: 20 },
//   input: { borderWidth: 1, marginVertical: 10, padding: 8 },
// });

// export default JoinVPN;


import React, { useEffect, useState } from 'react';
import { Button, Text, View } from 'react-native';
import useWebSocket from './signaling'; // Import WebSocket hook

const JoinVPN = () => {
  const { socket, connected } = useWebSocket();
  const [peerConnection, setPeerConnection] = useState(null);

  useEffect(() => {
    if (connected) {
      // Initialize WebRTC PeerConnection when connected to the signaling server
      const pc = new RTCPeerConnection();

      // Handle ICE candidate events
      pc.onicecandidate = (event) => {
        if (event.candidate) {
          console.log('Sending ICE candidate:', event.candidate);
          socket.send(JSON.stringify({ type: 'candidate', candidate: event.candidate }));
        }
      };

      setPeerConnection(pc);
    }

    return () => {
      // Clean up WebRTC connection on component unmount
      if (peerConnection) {
        peerConnection.close();
      }
    };
  }, [connected]);

  const handleOffer = async (offer) => {
    if (!peerConnection) return;

    try {
      await peerConnection.setRemoteDescription(new RTCSessionDescription(offer));

      const answer = await peerConnection.createAnswer();
      await peerConnection.setLocalDescription(answer);

      console.log('Sending answer:', answer);
      socket.send(JSON.stringify({ type: 'answer', answer }));
    } catch (error) {
      console.error('Error handling offer:', error);
    }
  };

  return (
    <View>
      <Text>Join VPN</Text>
      <Button title="Join VPN" onPress={() => {}} disabled={!connected} />
    </View>
  );
};

export default JoinVPN;

