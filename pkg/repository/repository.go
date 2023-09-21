package repository

import "database/sql"

type Authorization interface{}

type ToDoList interface{}

type ToDoItem interface{}

type Repository struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
