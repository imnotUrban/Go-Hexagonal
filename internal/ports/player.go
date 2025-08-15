package ports

import "api/internal/domain"

// PlayerService es una interfaz (Un contrato que no debe romperse) que define los métodos que cualquier servicio de jugador debe implementar.
type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
}
