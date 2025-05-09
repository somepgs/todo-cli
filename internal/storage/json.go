package storage

import (
	"encoding/json"
	"os"

	"github.com/somepgs/todo-cli/internal/entity"
)

type JSONStorage struct {
	FilePath string
}

func (s *JSONStorage) Load() ([]entity.Task, error) {
	file, err := os.Open(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []entity.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []entity.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *JSONStorage) Save(tasks []entity.Task) error {
	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

var _ Storage = (*JSONStorage)(nil)
