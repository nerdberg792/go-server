package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/nerdber792/go-server/routes"
)

func main() {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/Login", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true})
	})

	router.Run(":" + port)
}
