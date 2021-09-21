package mongodb

import (
	"context"
	"time"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// check if missing implementation
var _ repo.TodoRepo = new(todoRepo)

func NewTodoRepo(col *mongo.Collection) repo.TodoRepo {
	return &todoRepo{col}
}

type todoRepo struct {
	todosCollection *mongo.Collection
}

func (tr *todoRepo) GetTodoListOf(username string) []entity.Todo {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cursor, err := tr.todosCollection.Find(ctx, bson.M{"username": username})
	if err != nil {
		return nil
	}

	var todos []entity.Todo
	err = cursor.All(ctx, &todos)
	if err != nil {
		return nil
	}

	return todos
}

func (tr *todoRepo) GetTodoByID(id string) *entity.Todo {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var todo entity.Todo
	err := tr.todosCollection.FindOne(ctx, bson.M{"id": id}).Decode(&todo)
	if err != nil {
		return nil
	}

	return &todo
}

func (tr *todoRepo) CreateTodo(todo *entity.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := tr.todosCollection.InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (tr *todoRepo) DeleteTodo(todoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := tr.todosCollection.DeleteOne(ctx, bson.M{"id": todoID})
	if err != nil {
		return err
	}

	return nil
}

func (tr *todoRepo) UpdateTodo(todo *entity.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := tr.todosCollection.UpdateOne(
		ctx,
		bson.M{"id": todo.ID},
		bson.D{
			{
				"$set",
				bson.D{
					{Key: "content", Value: todo.Content},
					{Key: "isDone", Value: todo.IsDone},
				},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
