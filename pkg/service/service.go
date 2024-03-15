package service

import (
	GO_Project "github.com/zhetibayev0410/Go-1-project"
	"github.com/zhetibayev0410/Go-1-project/pkg/repository"
)

type Authorization interface {
	CreateUser(user GO_Project.User) (int error, err error)
}

type TodoList interface {
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
		//Authorization: NewAuthService(repos.Authorization),
	}
}
