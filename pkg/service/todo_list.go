package service

import (
	"awesomeProject"
	"awesomeProject/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func (t *TodoListService) Delete(listId int) error {
	return t.repo.Delete(listId)
}

func (t *TodoListService) GetById(listId int) (awesomeProject.TodoList, error) {
	return t.repo.GetById(listId)
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list awesomeProject.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]awesomeProject.TodoList, error) {
	return t.repo.GetAll(userId)
}
