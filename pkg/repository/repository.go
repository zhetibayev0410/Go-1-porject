package repository

import (
	"github.com/jmoiron/sqlx"
	GO_Project "github.com/zhetibayev0410/Go-1-project"
)

type Authorization interface {
	CreateUser(user GO_Project.User) (int, error)
	GetUser(username, password string) (GO_Project.User, error)
}

type TodoList interface {
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
	}
}
