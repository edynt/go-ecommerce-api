package controllers

import (
	"fmt"

	"github.com/edynt/go-ecommerce-api/internal/services"
	"github.com/edynt/go-ecommerce-api/internal/vo"
	"github.com/edynt/go-ecommerce-api/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params = vo.UserRegistrationRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid)
		return
	}

	fmt.Printf("Email params: %v", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
