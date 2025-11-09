package rest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HttpHandler struct {
	Router *gin.Engine
	DB     *gorm.DB
}
