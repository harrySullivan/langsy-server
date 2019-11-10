package controllers

import (
	"github.com/gin-gonic/gin"
)

type MainController struct {
	Controller
}

func (MainController) Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Langsy Server. Leave or DIE.",
	})
}

func (MainController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ashwat watches fifty shades of grey",
	})
}
