package app

import (
	"go-api/models"
	"go-api/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	// convert request body to Todo struct
	if err := c.BodyParser(&todo); err != nil {
		return c.Status((http.StatusBadRequest)).JSON(err.Error())
	}

	result, err := h.Service.TodoInsert(todo)

	if err != nil || result.Status == false {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(result)
}
