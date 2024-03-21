package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mohitmilindthakur/hotel-api/db"
	"github.com/mohitmilindthakur/hotel-api/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) == 0 {
		return errors.New("validation error")
	}
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userPayload types.CreateUserParams
	c.BodyParser(&userPayload)
	user, valid := types.NewUserFromParams(userPayload)
	if !valid {
		return errors.New("validation error")
	}
	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(*insertedUser)
}
