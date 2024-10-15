import { configureStore } from '@reduxjs/toolkit';
import vpnReducer from './vpnSlice';

export const store = configureStore({
  reducer: {
    vpn: vpnReducer,
  },
});
