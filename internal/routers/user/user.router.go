package user

import (
	"github.com/edynt/go-ecommerce-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authentication())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get-info")
	}
}
