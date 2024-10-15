import React from 'react';
import { View, Text, FlatList, StyleSheet } from 'react-native';

const PeerList = ({ peers }) => (
  <FlatList
    data={peers}
    keyExtractor={(item) => item.id}
    renderItem={({ item }) => (
      <View style={styles.peerItem}>
        <Text>{item.name}</Text>
      </View>
    )}
  />
);

const styles = StyleSheet.create({
  peerItem: { padding: 10, borderBottomWidth: 1 },
});

export default PeerList;
