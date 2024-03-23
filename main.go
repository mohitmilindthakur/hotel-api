package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohitmilindthakur/hotel-api/api"
	"github.com/mohitmilindthakur/hotel-api/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(400)
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("port", "5000", "Port number for the app")
	flag.Parse()
	port := ":" + *listenAddr

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, dbname))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user/:id", userHandler.GetUserById)
	apiv1.Put("/user/:id", userHandler.UpdateUser)
	apiv1.Get("/users", userHandler.GetUsers)
	apiv1.Post("/user", userHandler.CreateUser)
	apiv1.Delete("/user/:id", userHandler.DeleteUser)

	app.Listen(port)
}
