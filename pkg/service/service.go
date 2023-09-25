package service

import (
	structs "rest"
	"rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GenerateToken(userName, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list structs.TodoList) (int, error)
	GetAll(userId int) ([]structs.TodoList, error)
}

type ToDoItem interface{}

type Serivce struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewService(repo *repository.Repository) *Serivce {
	return &Serivce{
		Authorization: NewAuthService(repo.Authorization),
		ToDoList:      NewTodoListService(repo.ToDoList),
	}
}
