package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-dev-container/internal/database"
	"go-dev-container/pkg/compute"
)

func main() {
	time.Sleep(3 * time.Second)

	var a int = 4
	var b int = 9
	result := compute.SumNumber(a, b)

	log.Printf("Sum number %v, %v, result %v \n", a, b, result)

	if err := database.Initialize(); err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	database.Query()

	defer database.Close()

	// Initialize gin default instance
	router := gin.Default()

	// Create router group
	v1 := router.Group("/api/v1")

	// Register a route
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on
}
