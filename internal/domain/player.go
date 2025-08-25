package domain

import "time"

type Player struct {
	ID           interface{} `json:"id" bson:"_id"`
	FirstName    string      `json:"name" bson:"first_name" binding:"required"`
	CreationTime time.Time   `json:"creation_time" bson:"creation_time"`
}
