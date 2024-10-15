import React from 'react';
import { useVPN } from '../context/VPNContext';
import HostVPN from '../components/HostVPN';

const HostScreen = () => {
  const { dispatch } = useVPN();

  const handleCreateVPN = (sessionID) => {
    dispatch({ type: 'SET_SESSION_ID', payload: sessionID });
    dispatch({ type: 'SET_CONNECTION_STATUS', payload: 'connected' });
    console.log(`VPN created with session ID: ${sessionID}`);
  };

  return <HostVPN onCreate={handleCreateVPN} />;
};

export default HostScreen;
