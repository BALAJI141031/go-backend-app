package api

import (
	"backend-app/configs"
	"backend-app/internal/api/rest"
	"backend-app/internal/api/rest/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// starts with capital letter so publicly visible
func StartServer(config configs.AppConfig) {

	fmt.Println("Starting server on port: ", config.ServerPort)
	//create a gin router
	router := gin.Default()

	httpHandler := &rest.HttpHandler{
		Router: router,
	}

	setupRoutes(httpHandler)

	router.GET("health", healthCheckHandler)

	router.Run(config.ServerPort)
}

func setupRoutes(httpHandler *rest.HttpHandler) {
	//user handler
	handlers.SetupUserRoutes(httpHandler)

	//catalog handler
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Im breathing...",
	})
}
