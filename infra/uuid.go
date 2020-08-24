package infra

import (
	guuid "github.com/google/uuid"
	"gitlab.com/tsuchinaga/cli-todo/app/repository"
)

func NewUUID() repository.UUIDRepository {
	return &uuid{}
}

type uuid struct{}

func (u *uuid) New() (string, error) {
	uu, err := guuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uu.String(), nil
}
