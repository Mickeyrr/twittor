package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ObtenerTweets es la estructura con la que obtenemos los Tweets */
type ObtenerTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userID" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
