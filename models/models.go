package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Animation struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AnimationName   string             `json:"animationName,omitempty" bson:"animationName,omitempty"`
	Animationlottie string             `json:"animationlottie,omitempty" bson:"animationlottie,omitempty"`
}
