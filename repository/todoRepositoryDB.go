package repository

import (
	"context"
	"errors"
	"go-api/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	// will get data type of Todo
	Insert(todo models.Todo) (bool, error)
}

// Insert will insert a new todo
func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancel() will be called when the surrounding function returns
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("Error while inserting todo")
		return false, err
	}

	return true, nil
}

// NewTodoRepository will create an object that represent the TodoRepository interface
func NewTodoRepository(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
