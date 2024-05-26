module jmich.dev/todo-app/backend

go 1.22.2

require (
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/gorilla/mux v1.8.1
)

require github.com/rs/cors v1.11.0 // indirect

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	go.uber.org/atomic v1.11.0 // indirect
)

replace jmich.dev/todo-app/backend => /Users/jackm/Developer/todo-app/backend
