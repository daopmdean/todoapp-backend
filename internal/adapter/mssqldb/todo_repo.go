package mssqldb

import (
	"errors"
	"todoapp/internal/entity"
	"todoapp/internal/usecase/repo"

	"gorm.io/gorm"
)

// check if missing implementation
var _ repo.TodoRepo = new(todoRepo)

func NewTodoRepo(db *gorm.DB) repo.TodoRepo {
	return &todoRepo{db}
}

type todoRepo struct {
	db *gorm.DB
}

func (tr *todoRepo) GetTodoListOf(username string) []entity.Todo {
	var todos []entity.Todo
	tr.db.Where("username = ?", username).Find(&todos)
	return todos
}

func (tr *todoRepo) GetTodoByID(id string) *entity.Todo {
	var todo entity.Todo
	tr.db.Where("id = ?", id).Find(&todo)
	return &todo
}

func (tr *todoRepo) CreateTodo(todo *entity.Todo) error {
	result := tr.db.Create(todo)
	if result.Error != nil {
		return errors.New("cannot create todo with gorm: " + result.Error.Error())
	}
	return nil
}

func (tr *todoRepo) DeleteTodo(todoID string) error {
	result := tr.db.Delete(&entity.Todo{}, todoID)
	if result.Error != nil {
		return errors.New("cannot delete todo with gorm: " + result.Error.Error())
	}
	return nil
}

func (tr *todoRepo) UpdateTodo(todo *entity.Todo) error {
	result := tr.db.Save(todo)
	if result.Error != nil {
		return errors.New("cannot update todo with gorm: " + result.Error.Error())
	}
	return nil
}
