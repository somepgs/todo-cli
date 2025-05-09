package service

import (
	"errors"
	"time"

	"github.com/somepgs/todo-cli/internal/entity"
	"github.com/somepgs/todo-cli/internal/storage"
)

type TaskService struct {
	store storage.Storage
}

func NewTaskService(store storage.Storage) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) Add(title string) error {
	tasks, err := s.store.Load()
	if err != nil {
		return err
	}

	newTask := entity.Task{
		ID:        s.getNextID(tasks),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	return s.store.Save(tasks)
}

func (s *TaskService) List() ([]entity.Task, error) {
	return s.store.Load()
}

func (s *TaskService) Done(id int) error {
	tasks, err := s.store.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return s.store.Save(tasks)
		}
	}

	return errors.New("задача не найдена")
}

func (s *TaskService) Delete(id int) error {
	tasks, err := s.store.Load()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.store.Save(tasks)
		}
	}

	return errors.New("задача не найдена")
}

func (s *TaskService) getNextID(tasks []entity.Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
