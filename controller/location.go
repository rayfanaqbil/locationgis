package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/rayfanaqbil/locationgis/models"
	"github.com/rayfanaqbil/locationgis/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindNearestRoad godoc
// @Summary Find the nearest road
// @Description Given latitude and longitude, find the nearest road within a 5km radius
// @Accept json
// @Produce json
// @Param latitude query float64 true "Latitude"
// @Param longitude query float64 true "Longitude"
// @Success 200 {object} models.Road "Found the nearest road"
// @Failure 400 {object} models.ErrorResponse "Invalid input format"
// @Failure 404 {object} models.NotFoundResponse "No nearby road found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /findnearestroad [get]
func FindNearestRoad(c *fiber.Ctx) error {
	type Request struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Invalid input format",
		})
	}

	roadsCollection := config.Ulbimongoconn.Collection("jalan")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"geometry": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{request.Longitude, request.Latitude},
				},
				"$maxDistance": 5000, // Maximum distance of 5km
			},
		},
	}

	var road models.Road
	err := roadsCollection.FindOne(ctx, filter).Decode(&road)
	if err == mongo.ErrNoDocuments {
		return c.Status(http.StatusNotFound).JSON(models.NotFoundResponse{
			Message: "No nearby road found",
		})
	} else if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(road)
}
