package controllers

import (
	"net/http"

	"github.com/edynt/go-ecommerce-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController { // tra ve con tro
	return &UserController{
		userService: services.NewUserService(),
	} // &: dia chi con tro
}

// uc: user controller
// us: user service

func (uc *UserController) GetUserByID(c *gin.Context) {
	name := c.Param("name")
	uid := c.DefaultQuery("uid", "0000000")

	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"name":    name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "tringuyen"},
	})
}
