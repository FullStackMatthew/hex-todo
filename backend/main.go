package main

import (
	"database/sql"
	"log"
	"os"

	"example.com/hex-todo/backend/infrastructure/repository"
	"example.com/hex-todo/backend/interfaces"
	"example.com/hex-todo/backend/usecases"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	// Repository and Interactor setup
	todoRepository := repository.NewTodoRepository(db)
	todoInteractor := usecases.TodoInteractor{
		Repository: todoRepository,
	}

	// Handler setup
	handlers := interfaces.Handlers{
		TodoInteractor: todoInteractor,
	}

	// Setting up the router
	router := gin.Default()
	router.POST("/todos", handlers.AddTodoHandler)
	router.GET("/todos/:id", handlers.GetTodoHandler)
	router.GET("/todos", handlers.GetAllTodosHandler)
	router.PUT("/todos/:id", handlers.UpdateTodoHandler)
	router.DELETE("/todos/:id", handlers.DeleteTodoHandler)

	// Start server
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
