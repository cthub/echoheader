package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/ip", func(c *gin.Context) {
		c.Request.Header.Set("your_client_ip", c.ClientIP())
		c.JSON(200, c.Request.Header)
	})

	router.Run(":8080")
}
