package main

import (
	"GO_API/configs"
	"log"
	"GO_API/routes" 
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	// Set up Gin router
	router := gin.Default()

	// Use CORS middleware to allow requests from frontend (localhost:4200)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, 
		ExposeHeaders:     []string{"Content-Length"},
		AllowCredentials:  true, 
	}))

	client := configs.ConnectDB()

	if client == nil {
		log.Fatal("MongoDB connection failed")
	}

	routes.EmployeeRoute(router)

    router.Run("localhost:8080")
}