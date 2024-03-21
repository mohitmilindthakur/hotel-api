package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

type CreateUserParams struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
}

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName        string             `bson:"firstName" json:"firstName"`
	LastName         string             `bson:"lastName" json:"lastName"`
	Email            string             `bson:"email" json:"email"`
	encrptedPassword string             `bson:"encrptedPassword" json:"-"`
}

func (c *CreateUserParams) isValid() bool {
	errs := validate.Struct(c)
	return errs == nil
}

func NewUserFromParams(params CreateUserParams) (*User, bool) {
	if !params.isValid() {
		return nil, false
	}
	return &User{
		FirstName:        params.FirstName,
		LastName:         params.LastName,
		Email:            params.Email,
		encrptedPassword: params.Password,
	}, true
}
