import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL, // uses env var
});

export const test = "hi";

export default api;
