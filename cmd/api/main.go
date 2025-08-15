package main

import (
	"log"
	"os"

	"api/cmd/api/handlers/player"
	"api/internal/repositories/mongo"
	playerMongo "api/internal/repositories/mongo/player"
	playerService "api/internal/services/player"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()
	ginEngine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	playerRepo := playerMongo.Repository{
		Client: client,
	}

	playerSrv := playerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv, // Aquí debes inyectar la implementación concreta del servicio de jugador
	}

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	ginEngine.Run(":8080")
}
