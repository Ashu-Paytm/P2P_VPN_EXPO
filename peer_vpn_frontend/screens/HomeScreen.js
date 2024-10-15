import React from 'react';
import { View, Button, StyleSheet } from 'react-native';

const HomeScreen = ({ navigation }) => (
  <View style={styles.container}>
    <Button title="Host VPN" onPress={() => navigation.navigate('Host')} />
    <Button title="Join VPN" onPress={() => navigation.navigate('Join')} />
  </View>
);

const styles = StyleSheet.create({
  container: { flex: 1, justifyContent: 'center', padding: 20 },
});

export default HomeScreen;
