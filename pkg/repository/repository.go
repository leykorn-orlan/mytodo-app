package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/leykorn-orlan/todo-app"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorisation: NewAuthPostgres(db)}
}
