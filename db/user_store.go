package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/mohitmilindthakur/hotel-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) (*[]types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, any) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll:   c.Database(dbname).Collection(collName),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	return s.coll.Drop(ctx)
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

func (s *MongoUserStore) GetUsers(ctx context.Context) (*[]types.User, error) {
	var users []types.User
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID
	user.ID = id.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("empty Id")
	}
	objId, err := ToObjectID(id)
	if err != nil {
		return err
	}
	res, err := s.coll.DeleteOne(ctx, bson.M{"_id": objId})
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, id string, obj interface{}) error {
	if id == "" {
		return errors.New("empty id")
	}

	objId, err := ToObjectID(id)
	if err != nil {
		return err
	}

	fmt.Println("update", obj)
	res, err := s.coll.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": obj})

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("not found")
	}
	return nil
}
