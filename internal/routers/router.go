package routers

import (
	"github.com/edynt/go-ecommerce-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", controllers.NewPongController().Pong)
		v1.GET("/user/:name", controllers.NewUserController().GetUserByID)
	}

	return r
}
