package player

import (
	"api/internal/domain"
	"context"
	"fmt"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {

	collection := r.Client.Database("game").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		return nil, fmt.Errorf("error inserting player into MongoDB: %v", err)
	}

	return insertResult.InsertedID, nil
}
