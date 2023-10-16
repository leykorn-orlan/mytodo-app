package repositories

type Authorisation interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
