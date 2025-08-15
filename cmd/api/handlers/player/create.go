package player

import (
	"api/internal/domain"

	"github.com/gin-gonic/gin"
)

/*
* Debe hacer:
* - Validar el JSON recibido
* - Validar (lógica de negocio)
* - Consumir el servicio
* - Traducir el response
 */
func (h Handler) CreatePlayer(c *gin.Context) {
	var playerCreateParams domain.Player
	if err := c.BindJSON(&playerCreateParams); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	insertedId, err := h.PlayerService.Create(playerCreateParams)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create player"})
	}

	c.JSON(200, gin.H{"message": insertedId})
}
