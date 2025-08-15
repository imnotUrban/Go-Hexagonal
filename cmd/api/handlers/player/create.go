package player

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Player struct {
	Name         string    `json:"name" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"-"`
}

func CreatePlayer(c *gin.Context) {
	var player Player
	if err := c.BindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	player.CreationTime = time.Now().UTC()

	c.JSON(200, gin.H{"message": "Player created"})
}
