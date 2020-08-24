package service

import (
	"gitlab.com/tsuchinaga/cli-todo/app/entity"
	"gitlab.com/tsuchinaga/cli-todo/app/repository"
)

func NewTODO(store repository.TODOStore, uuid repository.UUIDRepository, clock repository.ClockRepository) TODO {
	return &todo{
		store: store,
		uuid:  uuid,
		clock: clock,
	}
}

type TODO interface {
	NewTODO(title string) (entity.TODO, error)
	Save(id string, todo entity.TODO)
	List() map[string]entity.TODO
	Delete(id string)
}

type todo struct {
	store repository.TODOStore
	uuid  repository.UUIDRepository
	clock repository.ClockRepository
}

func (s *todo) NewTODO(title string) (entity.TODO, error) {
	id, err := s.uuid.New()
	if err != nil {
		return entity.TODO{}, err
	}
	return entity.TODO{
		ID:          id,
		Title:       title,
		CreatedTime: s.clock.Now(),
	}, nil
}

func (s *todo) Save(id string, todo entity.TODO) {
	s.store.Set(id, todo)
}

func (s *todo) List() map[string]entity.TODO {
	return s.store.All()
}

func (s *todo) Delete(id string) {
	s.store.Delete(id)
}
