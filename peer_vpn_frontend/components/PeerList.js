import React from 'react';
import { View, Text, FlatList, StyleSheet, TouchableOpacity } from 'react-native';

// PeerList component to display a list of connected peers
const PeerList = ({ peers, onSelectPeer }) => {
  const renderItem = ({ item }) => (
    <TouchableOpacity style={styles.peerItem} onPress={() => onSelectPeer(item)}>
      <View style={styles.peerInfo}>
        <Text style={styles.peerName}>{item.name}</Text>
        <Text style={styles.peerStatus}>
          {item.isActive ? 'Active' : 'Inactive'}
        </Text>
      </View>
    </TouchableOpacity>
  );

  return (
    <FlatList
      data={peers}
      keyExtractor={(item) => item.id}
      renderItem={renderItem}
      ListEmptyComponent={<Text style={styles.emptyMessage}>No peers available</Text>}
    />
  );
};

const styles = StyleSheet.create({
  peerItem: {
    padding: 15,
    borderBottomWidth: 1,
    borderColor: '#ddd',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  peerInfo: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  peerName: {
    fontSize: 16,
    fontWeight: 'bold',
    marginRight: 10,
  },
  peerStatus: {
    fontSize: 14,
    color: '#555',
  },
  emptyMessage: {
    textAlign: 'center',
    fontSize: 16,
    color: '#888',
    padding: 20,
  },
});

export default PeerList;
