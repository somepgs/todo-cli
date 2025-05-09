package main

import (
	"fmt"
	"log"

	"github.com/somepgs/todo-cli/internal/service"
	"github.com/somepgs/todo-cli/internal/storage"
)

func main() {
	store := &storage.JSONStorage{FilePath: "tasks.json"}
	taskService := service.NewTaskService(store)

	if err := taskService.Add("Прочитать документацию по Go"); err != nil {
		log.Fatal(err)
	}

	if err := taskService.Add("Посмотреть вебинар по архитектуре"); err != nil {
		log.Fatal(err)
	}

	tasks, err := taskService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Все задачи:")
	for _, task := range tasks {
		fmt.Printf("[%d] %s (Выполнена: %v)\n", task.ID, task.Title, task.Done)
	}

	if err := taskService.Done(1); err != nil {
		log.Fatal(err)
	}

	tasks, err = taskService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nПосле выполнения задачи с ID=1:")
	for _, task := range tasks {
		fmt.Printf("[%d] %s (Выполнена: %v)\n", task.ID, task.Title, task.Done)
	}

	if err := taskService.Delete(2); err != nil {
		log.Fatal(err)
	}

	tasks, err = taskService.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nПосле удаления задачи с ID=2:")
	for _, task := range tasks {
		fmt.Printf("[%d] %s (Выполнена: %v)\n", task.ID, task.Title, task.Done)
	}
}
