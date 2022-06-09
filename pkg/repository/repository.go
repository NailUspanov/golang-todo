package repository

import (
	"awesomeProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user awesomeProject.User) (int, error)
	GetUser(username, password string) (awesomeProject.User, error)
}

type TodoList interface {
	Create(userId int, list awesomeProject.TodoList) (int, error)
	GetAll(userId int) ([]awesomeProject.TodoList, error)
	GetById(id int) (awesomeProject.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
