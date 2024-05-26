import React from "react";
import { Todo } from "../api";

interface TodoItemProps {
  todo: Todo;
}

const TodoItem: React.FC<TodoItemProps> = ({ todo }) => {
  return (
    <li>
      <h2>{todo.title}</h2>
      <p>Status: {todo.status}</p>
    </li>
  );
};

export default TodoItem;
