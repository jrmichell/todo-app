import React, { useState } from "react";
import { createTodo, Todo } from "../api";

interface TodoFormProps {
  onNewTodo: (todo: Todo) => void;
}

const TodoForm: React.FC<TodoFormProps> = ({ onNewTodo }) => {
  const [title, setTitle] = useState<string>("");
  const [status, setStatus] = useState<string>("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const newTodo = await createTodo({ title, status });
      console.log("New todo created:", newTodo);
      onNewTodo(newTodo);
      setTitle("");
      setStatus("");
    } catch (err) {
      console.error("Error creating todo:", err);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="title">Title:</label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      </div>
      <div>
        <label htmlFor="status">Status:</label>
        <input
          type="text"
          id="status"
          value={status}
          onChange={(e) => setStatus(e.target.value)}
        />
      </div>
      <button type="submit">Add Todo</button>
    </form>
  );
};

export default TodoForm;
