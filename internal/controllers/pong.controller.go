package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("-----> My handler")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...",
		"users":   []string{"cr7", "m10", "tringuyen"},
	})
}

func PongParams(c *gin.Context) {
	name := c.Param("name")
	uid := c.DefaultQuery("uid", "000000")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong params...",
		"name":    name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "tringuyen"},
	})
}
