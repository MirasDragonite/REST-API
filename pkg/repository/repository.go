package repository

import (
	"database/sql"

	structs "rest"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
}

type ToDoList interface{}

type ToDoItem interface{}

type Repository struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPost(db),
	}
}
