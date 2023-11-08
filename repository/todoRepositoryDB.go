package repository

import (
	"context"
	"errors"
	"go-api/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	// will get data type of Todo
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

// Insert will insert a new todo
func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancel() will be called when the surrounding function returns
	defer cancel()

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("Error inserting todo")
		return false, err
	}

	return true, nil
}

// GetAll will get all todos
func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todo []models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// bson.M{} will empty the filter and get all todos
	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		err := result.Decode(&todo)
		if err != nil {
			log.Fatalln(err)
		}
		// todo... means append each todo to todos
		todos = append(todos, todo...)
	}

	return todos, nil
}

// Delete will delete a todo
func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// bson.M{} will empty the filter and get all todos
	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil || result.DeletedCount <= 0 {
		log.Fatalln(err)
		return false, err
	}

	return true, nil
}

// NewTodoRepository will create an object that represent the TodoRepository interface
func NewTodoRepository(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
