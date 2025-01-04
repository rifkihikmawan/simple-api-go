package server

import (
	"log"
	"new-go-project/config"
	"new-go-project/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server interface {
	Start()
}

type server struct {
	app *echo.Echo
	db  *gorm.DB
}

func NewServer() Server {
	echoApp := echo.New()

	// Initialize database
	dbConfig := config.NewDatabase()
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database!")

	return &server{
		app: echoApp,
		db:  db,
	}
}

func (s *server) Start() {
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())

	log.Printf("Server is running on http://localhost:8080")

	// register API routes
	s.app.GET("/hello-world", s.handleHelloWorld)
	s.app.GET("/hello/:name", s.handleHello)

	// User routes
	s.app.GET("/users", s.handleGetAllUsers)
	s.app.POST("/users", s.handleCreateUser)
	s.app.GET("/users/:id", s.handleGetUserByID)
	s.app.PUT("/users/:id", s.handleUpdateUser)
	s.app.PUT("/users/:id/activate", s.handleActivateUser)
	s.app.DELETE("/users/:id", s.handleDeleteUser)

	s.app.Start(":8080")
}
