import React, { useState, useEffect } from "react";
import TodoList from "./components/TodoList";
import TodoForm from "./components/TodoForm";
import { Todo, getTodos } from "./api";
import "./App.css";

const App: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);

  const handleNewTodo = (newTodo: Todo) => {
    setTodos((prevTodos) => [...prevTodos, newTodo]);
  };

  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const todos = await getTodos();
        setTodos(todos);
      } catch (err) {
        console.error("Error fetching todos:", err);
      }
    };
    fetchTodos();
  }, []);

  return (
    <div className="App">
      <h1>Todo App</h1>
      <TodoForm onNewTodo={handleNewTodo} />
      <TodoList todos={todos} />
    </div>
  );
};

export default App;
