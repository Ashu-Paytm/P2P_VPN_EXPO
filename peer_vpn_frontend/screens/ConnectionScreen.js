import React from 'react';
import ConnectionStatus from '../components/ConnectionStatus';
import PeerList from '../components/PeerList';

const ConnectionScreen = () => {
  const peers = [{ id: '1', name: 'Peer 1' }, { id: '2', name: 'Peer 2' }];

  return (
    <>
      <ConnectionStatus status="Connected" />
      <PeerList peers={peers} />
    </>
  );
};

export default ConnectionScreen;
