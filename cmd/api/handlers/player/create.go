package player

import (
	"api/internal/domain"
	"errors"
	"net/http"

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
		RespondError(err, c)
	}

	c.JSON(200, gin.H{"message": insertedId})
}

func RespondError(err error, c *gin.Context) {
	var appErr *domain.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case domain.ErrorCodeDuplicateKey:
			c.JSON(http.StatusConflict, appErr)
			return
		case domain.ErrorCodeNotFound:
			c.JSON(http.StatusNotFound, appErr)
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}
}
