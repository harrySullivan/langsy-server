package controllers

import (
	"github.com/gin-gonic/gin"
)

// MainController is the Main handler for the client and pinger
type MainController struct {
	Controller
}

// Homepage is just a page to make fun of fools
func (MainController) Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Langsy Server. Leave or DIE.",
	})
}

// Ping is used for uptime
func (MainController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ashwat watches fifty shades of grey",
	})
}
