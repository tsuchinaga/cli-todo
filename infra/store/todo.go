package store

import (
	"gitlab.com/tsuchinaga/cli-todo/app/entity"
	"gitlab.com/tsuchinaga/cli-todo/app/repository"
)

func NewTODO() repository.TODOStore {
	return &todo{store: map[string]entity.TODO{}}
}

type todo struct {
	store map[string]entity.TODO
}

func (s *todo) All() map[string]entity.TODO {
	return s.store
}

func (s *todo) Set(key string, todo entity.TODO) {
	s.store[key] = todo
}

func (s *todo) Delete(key string) {
	delete(s.store, key)
}
