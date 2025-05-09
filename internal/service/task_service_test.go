package service

import (
	"testing"

	"github.com/somepgs/todo-cli/internal/entity"
)

type MockStorage struct {
	tasks []entity.Task
}

func (m *MockStorage) Load() ([]entity.Task, error) {
	return m.tasks, nil
}

func (m *MockStorage) Save(tasks []entity.Task) error {
	m.tasks = tasks
	return nil
}

func TestTaskService_Add(t *testing.T) {
	mockStore := &MockStorage{}
	svc := NewTaskService(mockStore)

	err := svc.Add("Новая задача")
	if err != nil {
		t.Fatalf("Ошибка добавления задачи: %v", err)
	}

	if len(mockStore.tasks) != 1 {
		t.Fatalf("Ожидалось 1 задача, получили %d", len(mockStore.tasks))
	}

	if mockStore.tasks[0].Title != "Новая задача" {
		t.Errorf("Название задачи неверное: %s", mockStore.tasks[0].Title)
	}
}

func TestTaskService_Done(t *testing.T) {
	mockStore := &MockStorage{
		tasks: []entity.Task{{ID: 1, Title: "Тестовая задача", Done: false}},
	}
	svc := NewTaskService(mockStore)

	err := svc.Done(1)
	if err != nil {
		t.Fatalf("Ошибка выполнения задачи: %v", err)
	}

	if !mockStore.tasks[0].Done {
		t.Errorf("Задача должна быть выполнена")
	}
}

func TestTaskService_Delete(t *testing.T) {
	mockStore := &MockStorage{
		tasks: []entity.Task{
			{ID: 1, Title: "Тестовая задача", Done: false},
			{ID: 2, Title: "Вторая задача", Done: false},
		},
	}
	svc := NewTaskService(mockStore)

	err := svc.Delete(1)
	if err != nil {
		t.Fatalf("Ошибка удаления задачи: %v", err)
	}

	if len(mockStore.tasks) != 1 {
		t.Fatalf("Ожидалась 1 задача после удаления, получили %d", len(mockStore.tasks))
	}

	if mockStore.tasks[0].ID != 2 {
		t.Errorf("Осталась неверная задача: %+v", mockStore.tasks[0])
	}
}
