package service

import "rest/pkg/repository"

type Authorization interface{}

type ToDoList interface{}

type ToDoItem interface{}

type Serivce struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewService(repo *repository.Repository) *Serivce {
	return &Serivce{}
}
