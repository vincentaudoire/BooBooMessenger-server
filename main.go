package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/vincentaudoire/BooBooMessenger-server/model"
	"github.com/vincentaudoire/BooBooMessenger-server/repository"
	"github.com/vincentaudoire/BooBooMessenger-server/rest"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

var messages []model.Message
var db *sql.DB

func main() {

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// Setting up the database
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Failed to connect to database")
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println("Failed to create database migration driver")
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Println("Failed to create migration")
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Println("Failed to execute migration")
			log.Fatal(err)
		}
	}

	messageRepository := repository.New(db)

	messageController := rest.NewMessageController(messageRepository)

	// Setting up the router
	router := gin.Default()
	router.GET("/messages", messageController.GetAllMessage)
	router.PUT("messages/:id/printed", messageController.MarkMessageAsRead)
	router.POST("/messages", messageController.SaveNewMessage)
	router.Run() // listen and serve on 0.0.0.0:8080
}
