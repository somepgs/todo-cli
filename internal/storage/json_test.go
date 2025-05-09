package storage

import (
	"os"
	"testing"
	"time"

	"github.com/somepgs/todo-cli/internal/entity"
)

func TestJSONStorage_SaveLoad(t *testing.T) {
	filepath := "test_tasks.json"
	defer os.Remove(filepath)

	store := &JSONStorage{FilePath: filepath}

	tasksToSave := []entity.Task{
		{ID: 1, Title: "Test task 1", Done: false, CreatedAt: time.Now()},
		{ID: 2, Title: "Test task 2", Done: true, CreatedAt: time.Now()},
	}

	if err := store.Save(tasksToSave); err != nil {
		t.Fatalf("Ошибка сохранения задач: %v", err)
	}

	loadedTasks, err := store.Load()
	if err != nil {
		t.Fatalf("Ошибка загрузки задач: %v", err)
	}

	if len(loadedTasks) != 2 {
		t.Errorf("Ожидалось 2 задачи, получили %d", len(loadedTasks))
	}

	for i, task := range loadedTasks {
		if task.Title != tasksToSave[i].Title || task.Done != tasksToSave[i].Done {
			t.Errorf("Несоответствие задач: ожидалось %+v, получили %+v", tasksToSave[i], task)
		}
	}
}
