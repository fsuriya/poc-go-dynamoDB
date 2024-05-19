package routers

import (
	"poc-go-dynamodb/routers/auth"
	"poc-go-dynamodb/routers/kol"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define auth route group
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signin", auth.SignIn)
		authGroup.POST("/register", auth.Register)
	}

	// Define KOL route group
	kolGroup := r.Group("/kol")
	{
		kolGroup.POST("/", kol.CreateKOL)
		kolGroup.PATCH("/:kol_id", kol.UpdateKOL)
		kolGroup.DELETE("/:kol_id", kol.DeleteKOL)
		kolGroup.GET("/", kol.GetKOL)
	}

	return r
}
