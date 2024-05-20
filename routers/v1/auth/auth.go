package auth

import (
	"net/http"

	"poc-go-dynamodb/pkg/app"
	"poc-go-dynamodb/service/auth_service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func SignIn(c *gin.Context) {
	// Handle sign-in logic here (e.g., validate credentials)
	c.JSON(http.StatusOK, gin.H{"message": "Sign-in successful (placeholder)"})
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var a auth
	if err := c.BindJSON(&a); err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		appG.Response(http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return
	}

	existsUser, err := auth_service.ExistsUser(a.Username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}
	println(existsUser)

	// Handle registration logic here (e.g., create new user)
	c.JSON(http.StatusOK, gin.H{"message": "Registration " + a.Username + " successful"})
}
