package handlers

import (
	"backend-app/internal/api/rest"
	"backend-app/internal/dto"
	"backend-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// user service
	userService service.UserService
}

func SetupUserRoutes(httpHandler *rest.HttpHandler) {

	userService := service.UserService{}
	userHandler := UserHandler{userService: userService}

	// httpHandler.Router.GET("/users", getAllUsers)
	httpHandler.Router.POST("/register", userHandler.Register)
	httpHandler.Router.POST("user/login", userHandler.Login)

	httpHandler.Router.GET("verificationCode", userHandler.GetVerificationCode)
	httpHandler.Router.POST("verificationCode", userHandler.VerifyCode)
	httpHandler.Router.GET("user/profile", userHandler.Profile)

	httpHandler.Router.POST("cart", userHandler.AddToCart)
	httpHandler.Router.GET("cart", userHandler.GetCart)
	httpHandler.Router.DELETE("cart", userHandler.ClearCart)
}

func (userHandler *UserHandler) Register(c *gin.Context) {
	userDto := dto.UserSignupRequestDto{}
	err := c.ShouldBindJSON(&userDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := userHandler.userService.SignupUser(userDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User is Registered", "token": token})
}

func (UserHandler *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User is Logged In"})
}

func (userHandler *UserHandler) GetVerificationCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get Verification Code"})
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
