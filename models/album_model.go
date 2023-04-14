package models

import "go.mongodb.org/mongo-driver/bson/primitive"
// album represents data about a record album.
type Album struct {
    ID     primitive.ObjectID  `json:"id,omitempty"`
    Title  string  `json:"title,omitempty" validate:"required"`
    Artist string  `json:"artist,omitempty" validate:"required"`
    Price  float64 `json:"price" validate:"required"`
}

