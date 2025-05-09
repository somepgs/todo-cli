package main

import (
	"fmt"
	"log"
	"time"

	"github.com/somepgs/todo-cli/internal/entity"
	"github.com/somepgs/todo-cli/internal/storage"
)

func main() {
	store := &storage.JSONStorage{FilePath: "tasks.json"}

	task := entity.Task{
		ID:        1,
		Title:     "Написать тестовое приложение",
		Done:      false,
		CreatedAt: time.Now(),
	}

	err := store.Save([]entity.Task{task})
	if err != nil {
		log.Fatal("Ошибка сохранения:", err)
	}

	tasks, err := store.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки:", err)
	}

	fmt.Println("Загруженные задачи:", tasks)
}
