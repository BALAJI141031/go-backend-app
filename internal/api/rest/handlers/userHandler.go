package handlers

import (
	"backend-app/internal/api/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// user service

}

func SetupUserRoutes(httpHandler *rest.HttpHandler) {

	userHandler := UserHandler{}

	// httpHandler.Router.GET("/users", getAllUsers)
	httpHandler.Router.GET("/register", userHandler.Register)
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User is Registered"})
}
