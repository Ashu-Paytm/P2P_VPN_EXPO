import React from 'react';
import JoinVPN from '../components/JoinVPN';

const JoinScreen = () => {
  const handleJoinVPN = (sessionID) => {
    console.log(`Joining VPN with session ID: ${sessionID}`);
    // Add backend integration here
  };

  return <JoinVPN onJoin={handleJoinVPN} />;
};

export default JoinScreen;
