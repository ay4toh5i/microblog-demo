import axios from 'axios';
import applyCaseMiddleware from 'axios-case-converter';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

export const axiosApi = applyCaseMiddleware(api);
