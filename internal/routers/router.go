package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", Pong)
		v1.GET("/ping-param/:name", PongParams)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong 1111",
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
