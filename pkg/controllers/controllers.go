package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api/pkg/models"
	"golang-rest-api/pkg/utils"
	"log"
)

func RegisterUser(c *fiber.Ctx) error {
	// Connect DB
	dbUri := utils.LoadDotEnv("DB_URI")
	col := utils.ConnectCollection(dbUri, "blgy-app", "users")

	// Import Model
	p := new(models.Person)

	// Get Data
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(400)
	}

	// Store Data
	data := bson.D{{"name", p.Name}, {"surname", p.Surname}, {"email", p.Email}, {"password", p.Password}}

	// Send Data
	_, err := col.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}

	// Return nil
	return nil
}
func LoginUser(c *fiber.Ctx) error {

	// Connect DB
	dbUri := utils.LoadDotEnv("DB_URI")
	col := utils.ConnectCollection(dbUri, "blgy-app", "users")

	// Import Model
	p := new(models.Person)

	// Get Data
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(400)
	}

	// Store Data
	data := bson.D{{"email", p.Email}, {"password", p.Password}}

	// Response Data
	var result bson.M

	// Send Data
	err := col.FindOne(context.TODO(), data).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return c.SendStatus(400)
		}
		panic(err)
	}

	log.Println(result)

	// Return nil
	return nil
}
