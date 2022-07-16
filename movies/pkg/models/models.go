package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie is used to represent movie profile data
type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Rating      float64            `bson:"rating,omitempty"`
	CreatedOn   time.Time          `bson:"createdon,omitempty"`
	ImdbUrl     string             `bson:"imdb_url,omitempty"`
}
