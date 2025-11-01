package api

import (
	"backend-app/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

// starts with capital letter so publicly visible
func StartServer(config configs.AppConfig) {

	fmt.Println("Starting server on port: ", config.ServerPort)
	//create a gin router
	router := gin.Default()

	router.Run(config.ServerPort)
}
