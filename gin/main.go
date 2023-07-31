package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	// Ping test
	r.GET("/notify", func(c *gin.Context) {
		fmt.Println("---------------->callback")
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "success",
			"data": gin.H{
				"content": "pong",
			},
		})
	})
	r.GET("/return", func(c *gin.Context) {
		fmt.Println("---------------->return_url")
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "re",
			"data": gin.H{
				"content": "pong",
			},
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8999")
}
