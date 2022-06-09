package service

import (
	"awesomeProject"
	"awesomeProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user awesomeProject.User) (int, error)
	GenerateJWT(username, password string) (string, error)
	ParseToken(s string) (int, error)
}

type TodoList interface {
	Create(userId int, list awesomeProject.TodoList) (int, error)
	GetAll(userId int) ([]awesomeProject.TodoList, error)
	GetById(listId int) (awesomeProject.TodoList, error)
	Delete(listId int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
