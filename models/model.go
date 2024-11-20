package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Road struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Type       string             `bson:"type"`
	Geometry   Geometry           `bson:"geometry"`
	Properties Properties         `bson:"properties"`
}

type Geometry struct {
	Type        string        `bson:"type"`
	Coordinates [][]float64   `bson:"coordinates"`
}

type Properties struct {
	OsmID    int    `bson:"osm_id"`
	Name     string `bson:"name"`
	Highway  string `bson:"highway,omitempty"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type NotFoundResponse struct {
    Message string `json:"message"`
}