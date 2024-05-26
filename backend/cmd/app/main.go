package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"jmich.dev/todo-app/backend/cmd/internal/database"
	"jmich.dev/todo-app/backend/cmd/internal/models"
	"jmich.dev/todo-app/backend/cmd/internal/routes"
)

func main() {
	var resetDB bool
	flag.BoolVar(&resetDB, "reset-db", false, "Reset the database")
	flag.Parse()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	fmt.Printf("DATABASE_URL: %s\n", connectionString)

	if resetDB {
		err = database.ResetDatabase(connectionString)
		if err != nil {
			log.Fatalf("Error resetting database: %v", err)
		}
		fmt.Println("Database reset successfully")
		return
	}

	err = database.InitDB(connectionString)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Set the database for models
	models.SetDatabase(database.DB)

	err = database.RunMigrations(connectionString)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// Add CORS support
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
