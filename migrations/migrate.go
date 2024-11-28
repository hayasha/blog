package main

import (
	"github.com/hayasha/blog/initializers"
	"github.com/hayasha/blog/models"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToMysql()
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration successful")
}
