package player

import (
	"api/internal/domain"
	"errors"
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
		// Si el repo envuelve el error de clave duplicada con domain.ErrDuplicateKey
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Printf("Duplicate key error: %v", err)
			// Devuelve el AppError como error (puntero) y encadénalo
			return nil, fmt.Errorf("%w", &domain.AppError{
				Code:    domain.ErrorCodeDuplicateKey,
				Message: "A player with the same key already exists.",
			})
		}
		// Error genérico
		return nil, fmt.Errorf("error saving player: %w", err)
	}

	return insertedId, nil
}
