import React, { useState } from 'react';
import { View, Text, TextInput, Button, StyleSheet } from 'react-native';

const HostVPN = ({ onCreate }) => {
  const [sessionID, setSessionID] = useState('');

  return (
    <View style={styles.container}>
      <Text>Host a VPN Session</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter Session ID"
        value={sessionID}
        onChangeText={setSessionID}
      />
      <Button title="Create VPN" onPress={() => onCreate(sessionID)} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: { padding: 20 },
  input: { borderWidth: 1, marginVertical: 10, padding: 8 },
});

export default HostVPN;
