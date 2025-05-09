package storage

import "github.com/somepgs/todo-cli/internal/entity"

type Storage interface {
	Load() ([]entity.Task, error)
	Save([]entity.Task) error
}
