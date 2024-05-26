package routes

import (
	"github.com/gorilla/mux"
	"jmich.dev/todo-app/backend/cmd/internal/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
}
