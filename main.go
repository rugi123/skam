package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.FileAttachment("stepapidr.txt", "neskam.txt")
	})

	telegrambot.runBot()

	router.Run(":8080")
}
