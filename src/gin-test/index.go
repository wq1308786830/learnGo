package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5757") // listen and serve on 0.0.0.0:8080
}
