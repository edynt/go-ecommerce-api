package routers

import (
	"fmt"

	"github.com/edynt/go-ecommerce-api/internal/controllers"
	"github.com/edynt/go-ecommerce-api/internal/middlewars"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewars.AuthenMiddleware(), AA(), BB(), CC)

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", controllers.NewPongController().Pong)
		v1.GET("/user/:name", controllers.NewUserController().GetUserByID)
	}

	return r
}
