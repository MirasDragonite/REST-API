package repository

type Authorization interface{}

type ToDoList interface{}

type ToDoItem interface{}

type Repository struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewRepository() *Repository {
	return &Repository{}
}
