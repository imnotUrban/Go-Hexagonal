package player

import (
	"api/internal/domain"
	"fmt"
	"log"
	"time"
)

/*
* Lógica de negocio para crear un jugador.
 */
func (s Service) Create(player domain.Player) (interface{}, error) {
	player.CreationTime = time.Now().UTC()

	insertedId, err := s.Repo.Insert(player)
	if err != nil {
		log.Printf("Error saving player: %v", err)
		return nil, fmt.Errorf("error saving player: %v", err)
	}

	return insertedId, nil
}
