package player

import (
	"api/internal/ports"
)

type Service struct {
	Repo ports.PlayerRepository
}
