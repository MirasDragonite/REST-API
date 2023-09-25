package repository

import (
	"database/sql"

	structs "rest"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GetUser(username, password string) (structs.User, error)
}

type ToDoList interface {
	Create(userId int, list structs.TodoList) (int, error)
	GetAll(userId int) ([]structs.TodoList, error)
}

type ToDoItem interface{}

type Repository struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPost(db),
		ToDoList:      NewToDoListPost(db),
	}
}
