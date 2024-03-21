package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mohitmilindthakur/hotel-api/db"
	"github.com/mohitmilindthakur/hotel-api/types"
	"go.mongodb.org/mongo-driver/mongo"
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]interface{}{
				"success": false,
				"message": "Not found",
			})
		}
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

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errors.New("empty ID")
	}
	err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]bool{"success": true})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errors.New("empty id")
	}
	var user interface{}
	c.BodyParser(&user)
	err := h.userStore.UpdateUser(c.Context(), id, user)
	if err != nil {
		return err
	}
	return c.JSON(map[string]bool{"success": true})
}
