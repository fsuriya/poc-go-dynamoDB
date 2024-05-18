package main

import (
	"fmt"

	"poc-go-dynamodb/auth"
	"poc-go-dynamodb/kol"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define auth route group
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signin", auth.SignIn)
		authGroup.POST("/register", auth.Register)
	}

	// Define KOL route group
	kolGroup := router.Group("/kol")
	{
		kolGroup.POST("/", kol.CreateKOL)
		kolGroup.PATCH("/:kol_id", kol.UpdateKOL)
		kolGroup.DELETE("/:kol_id", kol.DeleteKOL)
		kolGroup.GET("/", kol.GetKOL)
	}

	// Start the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
