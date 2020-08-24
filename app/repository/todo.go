package repository

import "gitlab.com/tsuchinaga/cli-todo/app/entity"

type TODOStore interface {
	All() map[string]entity.TODO
	Set(key string, todo entity.TODO)
	Delete(key string)
}
