package main

import (
	"go-api/app"
	"go-api/configs"
	"go-api/repository"
	"go-api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create fiber app
	appRoute := fiber.New()

	// connect to database
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	// create repository
	TodoRepositoryDB := repository.NewTodoRepository(dbClient)

	// create service
	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	// create route
	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todo", td.GetAll)

	// run app
	port := configs.EnvPort()
	appRoute.Listen(":" + port)
}
