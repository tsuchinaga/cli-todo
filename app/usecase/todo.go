package usecase

import (
	"sort"

	"gitlab.com/tsuchinaga/cli-todo/app/entity"
	"gitlab.com/tsuchinaga/cli-todo/app/service"
)

func NewTODO(todoService service.TODO) TODO {
	return &todo{
		todoService: todoService,
	}
}

type TODO interface {
	List() []entity.TODO
	Create(title string) (string, error)
	Delete(id string) error
}

type todo struct {
	todoService service.TODO
}

func (u *todo) List() []entity.TODO {
	list := make([]entity.TODO, 0)
	for _, t := range u.todoService.List() {
		list = append(list, t)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedTime.Before(list[j].CreatedTime)
	})

	return list
}

func (u *todo) Create(title string) (string, error) {
	todo, err := u.todoService.NewTODO(title)
	if err != nil {
		return "", err
	}
	u.todoService.Save(todo.ID, todo)
	return todo.ID, nil
}

func (u *todo) Delete(id string) error {
	u.todoService.Delete(id)
	return nil
}
