import { createSlice } from '@reduxjs/toolkit';

const vpnSlice = createSlice({
  name: 'vpn',
  initialState: { sessionID: '', peers: [] },
  reducers: {
    setSessionID: (state, action) => {
      state.sessionID = action.payload;
    },
    addPeer: (state, action) => {
      state.peers.push(action.payload);
    },
  },
});

export const { setSessionID, addPeer } = vpnSlice.actions;
export default vpnSlice.reducer;
