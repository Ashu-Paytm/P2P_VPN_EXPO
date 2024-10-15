import axios from 'axios';
import { BASE_URL } from '../utils/constants';

// Create an Axios instance with default settings
const api = axios.create({
  baseURL: BASE_URL,
  timeout: 10000,
});

export const createSession = async (sessionID) => {
  try {
    const response = await api.post('/session/create', { sessionID });
    return response.data;
  } catch (error) {
    console.error('Error creating session:', error);
    throw error;
  }
};

export const joinSession = async (sessionID) => {
  try {
    const response = await api.post('/session/join', { sessionID });
    return response.data;
  } catch (error) {
    console.error('Error joining session:', error);
    throw error;
  }
};

export const getPeers = async (sessionID) => {
  try {
    const response = await api.get(`/session/${sessionID}/peers`);
    return response.data;
  } catch (error) {
    console.error('Error fetching peers:', error);
    throw error;
  }
};
