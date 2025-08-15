package player

import "api/internal/ports"

type Handler struct {
	PlayerService ports.PlayerService
}
