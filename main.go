package main

import (
	"github.com/gin-gonic/gin"
	"golang_workshop/facade"
	"net/http"
)

type ReverseAPIInput struct {
	Text string
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/reverse", func(c *gin.Context) {
		var apiInput ReverseAPIInput
		if err := c.BindJSON(&apiInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"reverse": "",
				"message": "Invalid input",
			})
		}
		rev := facade.Reverse(apiInput.Text)
		c.JSON(http.StatusOK, gin.H{
			"reverse": rev,
			"message": "",
		})
	})
	err := router.Run()
	if err != nil {
		return
	}
}
