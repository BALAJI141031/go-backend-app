package api

import (
	"backend-app/configs"
	"backend-app/internal/api/rest"
	"backend-app/internal/api/rest/handlers"
	"backend-app/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// starts with capital letter so publicly visible
func StartServer(config configs.AppConfig) {

	fmt.Println("Starting server on port: ", config.ServerPort)
	//create a gin router
	router := gin.Default()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	fmt.Println("Database URL: ", config.Dsn)
	if err != nil {
		panic("failed to connect database")
	}

	//run the migrations
	db.AutoMigrate(&domain.User{})

	httpHandler := &rest.HttpHandler{
		Router: router,
		DB:     db,
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
