package player

import (
	"api/internal/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {

	collection := r.Client.Database("game").Collection("players")
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("%w: error inserting player into MongoDB: %v", domain.ErrDuplicateKey, err.Error())
		}
		return nil, fmt.Errorf("error inserting player into MongoDB: %v", err)
	}

	return insertResult.InsertedID, nil
}
