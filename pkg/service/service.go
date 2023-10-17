package service

import "github.com/leykorn-orlan/todo-app/pkg/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
