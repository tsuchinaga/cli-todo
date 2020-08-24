package adapter

import (
	"gitlab.com/tsuchinaga/cli-todo/app/entity"
	"gitlab.com/tsuchinaga/cli-todo/app/usecase"
)

func NewTODO(todoUseCase usecase.TODO) TODO {
	return &todo{
		todoUseCase: todoUseCase,
	}
}

type TODO interface {
	List() ([]entity.TODO, error)
	Create(title string) (string, error)
	Delete(id string) error
}

type todo struct {
	todoUseCase usecase.TODO
}

func (a *todo) List() ([]entity.TODO, error) {
	return a.todoUseCase.List(), nil
}

func (a *todo) Create(title string) (string, error) {
	return a.todoUseCase.Create(title)
}

func (a *todo) Delete(id string) error {
	return a.todoUseCase.Delete(id)
}
