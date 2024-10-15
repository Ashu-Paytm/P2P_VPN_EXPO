import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const ConnectionStatus = ({ status }) => (
  <View style={styles.container}>
    <Text>Connection Status: {status}</Text>
  </View>
);

const styles = StyleSheet.create({
  container: { padding: 10, marginTop: 10 },
});

export default ConnectionStatus;
