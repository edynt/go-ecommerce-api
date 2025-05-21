package controllers

import (
	"github.com/edynt/go-ecommerce-api/internal/services"
	"github.com/edynt/go-ecommerce-api/response"
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
	response.SuccessResponse(c, 2001, uc.userService.GetInfoUser())
}
