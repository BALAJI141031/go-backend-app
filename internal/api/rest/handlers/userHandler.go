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
	httpHandler.Router.POST("/register", userHandler.Register)
	httpHandler.Router.POST("user/login", userHandler.Login)

	httpHandler.Router.GET("verifyCode", userHandler.GetVerifyCode)
	httpHandler.Router.POST("verifyCode", userHandler.VerifyCode)
	httpHandler.Router.GET("user/profile", userHandler.Profile)

	httpHandler.Router.POST("cart", userHandler.AddToCart)
	httpHandler.Router.GET("cart", userHandler.GetCart)
	httpHandler.Router.DELETE("cart", userHandler.ClearCart)
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User is Registered"})
}

func (UserHandler *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User is Logged In"})
}

func (userHandler *UserHandler) GetVerifyCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get Verify Code"})
}

func (userHandler *UserHandler) VerifyCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Code is Verified"})
}

func (userHandler *UserHandler) Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User Profile Fetched"})
}

func (userHandler *UserHandler) AddToCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Item Added to Cart"})
}

func (userHandler *UserHandler) GetCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Cart Fetched"})
}

func (userHandler *UserHandler) ClearCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Cart is cleared"})
}
