package api

import (
	"context"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mohitmilindthakur/hotel-api/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testmongouri = "mongodb://localhost:27017"
const dbname = "hotel-reservation-test"

type testdb struct {
	UserStore db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	tdb.UserStore.Drop(context.TODO())
}
func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmongouri))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}
func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	UserHandler := NewUserHandler(tdb.UserStore)
	app := fiber.New()
	app.Post("/", UserHandler.CreateUser)
	// httptest.NewRequest("POST", "/", )
	// t.Fail()
}
