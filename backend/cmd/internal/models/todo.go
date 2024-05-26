package models

import (
	"database/sql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}

func GetAllTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, status FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func CreateTodo(todo *Todo) error {
	err := db.QueryRow("INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id", todo.Title, todo.Status).Scan(&todo.ID)
	if err != nil {
		return err
	}
	return nil
}
