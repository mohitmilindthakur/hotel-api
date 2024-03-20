package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mohitmilindthakur/hotel-api/db"
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
	ctx := context.Background()

	fmt.Println(id)
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON([]int{1, 2, 3, 4, 5})
}
