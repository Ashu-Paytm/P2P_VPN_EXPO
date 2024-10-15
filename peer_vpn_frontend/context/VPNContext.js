import React, { createContext, useReducer, useContext } from 'react';

// Initial state
const initialState = {
  sessionID: '',
  peers: [],
  connectionStatus: 'disconnected',
};

// Reducer to manage state updates
const vpnReducer = (state, action) => {
  switch (action.type) {
    case 'SET_SESSION_ID':
      return { ...state, sessionID: action.payload };
    case 'ADD_PEER':
      return { ...state, peers: [...state.peers, action.payload] };
    case 'SET_CONNECTION_STATUS':
      return { ...state, connectionStatus: action.payload };
    default:
      return state;
  }
};

// Create context
const VPNContext = createContext();

// Provider component
export const VPNProvider = ({ children }) => {
  const [state, dispatch] = useReducer(vpnReducer, initialState);

  return (
    <VPNContext.Provider value={{ state, dispatch }}>
      {children}
    </VPNContext.Provider>
  );
};

// Custom hook to use the VPN context
export const useVPN = () => useContext(VPNContext);
