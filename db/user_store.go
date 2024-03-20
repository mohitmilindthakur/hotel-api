package db

import (
	"context"
	"fmt"

	"github.com/mohitmilindthakur/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll:   c.Database(DBNAME).Collection(collName),
	}
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	objectId, err := ToObjectID(id)
	if err != nil {
		return nil, err
	}
	err = s.coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
