import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL || "http://localhost:8080",
});

export interface Todo {
  id: number;
  title: string;
  status: string;
}

export const getTodos = async (): Promise<Todo[]> => {
  const response = await api.get("/todos");
  if (!response.data || !Array.isArray(response.data)) {
    throw new Error("Invalid response from server");
  }
  return response.data;
};

export const createTodo = async (todo: Omit<Todo, "id">): Promise<Todo> => {
  const response = await api.post("/todos", todo);
  return response.data;
};

export default api;
