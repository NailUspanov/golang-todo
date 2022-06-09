package repository

import (
	"awesomeProject"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func (t *TodoListPostgres) GetById(id int) (awesomeProject.TodoList, error) {
	var list awesomeProject.TodoList
	getByIdQuery := fmt.Sprintf("SELECT tl.* FROM %s tl WHERE tl.id = $1", todoListsTable)
	err := t.db.Get(&list, getByIdQuery, id)

	return list, err
}

func (t *TodoListPostgres) Create(userId int, list awesomeProject.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) GetAll(userId int) ([]awesomeProject.TodoList, error) {

	var lists []awesomeProject.TodoList
	getAllQuery := fmt.Sprintf("SELECT ls.* FROM %s ls INNER JOIN %s usr on ls.id = usr.list_id WHERE usr.user_id = $1", todoListsTable, usersListsTable)
	err := t.db.Select(&lists, getAllQuery, userId)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
