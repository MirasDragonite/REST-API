package service

import (
	structs "rest"
	"rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GenerateToken(userName, password string) (string, error)
}

type ToDoList interface{}

type ToDoItem interface{}

type Serivce struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewService(repo *repository.Repository) *Serivce {
	return &Serivce{Authorization: NewAuthService(repo.Authorization)}
}
