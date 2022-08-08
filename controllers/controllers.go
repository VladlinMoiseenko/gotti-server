package controllers

import (
	"context"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/VladlinMoiseenko/gotti/db"
	"github.com/VladlinMoiseenko/gotti/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllanimation(c *fiber.Ctx) error {
	animationCollection := db.MI.DB.Collection("gotti")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var gotti []models.Animation

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"animationName": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"animationlottie": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := animationCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := animationCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Animation Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var animation models.Animation
		cursor.Decode(&animation)
		gotti = append(gotti, animation)
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      gotti,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}

func GetAnimation(c *fiber.Ctx) error {
	animationCollection := db.MI.DB.Collection("gotti")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var animation models.Animation
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := animationCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Animation Not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&animation)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Animation Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    animation,
		"success": true,
	})
}

func AddAnimation(c *fiber.Ctx) error {
	animationCollection := db.MI.DB.Collection("gotti")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	animation := new(models.Animation)

	if err := c.BodyParser(animation); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	//log.Println("c = ", animation.Animationlottie)

	result, err := animationCollection.InsertOne(ctx, animation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Animation failed to insert",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Animation inserted successfully",
	})

}
