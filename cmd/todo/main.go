package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/somepgs/todo-cli/internal/service"
	"github.com/somepgs/todo-cli/internal/storage"
)

const usage = `Ипользование:
	todo add "название задачи" - добавить задачу
	todo list				   - список задач
	todo done <id>			   - завершить задачу
	todo delete <id>		   - удалить задачу`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	store := &storage.JSONStorage{FilePath: "tasks.json"}
	taskService := service.NewTaskService(store)

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		addCmd.Parse(os.Args[2:])
		title := addCmd.Arg(0)
		if title == "" {
			log.Fatal("Не указано название задачи")
		}
		if err := taskService.Add(title); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Задача добавлена!")

	case "list":
		tasks, err := taskService.List()
		if err != nil {
			log.Fatal(err)
		}
		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "X"
			}
			fmt.Printf("[%s] %-3d %-30s (%s)\n", status, task.ID, task.Title, task.CreatedAt.Format("2006-01-02 15:04"))
		}

	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		doneCmd.Parse(os.Args[2:])
		id, err := strconv.Atoi(doneCmd.Arg(0))
		if err != nil {
			log.Fatal("Неправильно указан id задачи")
		}
		if err := taskService.Done(id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Задача %d отмечена выполненной!\n", id)

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		delCmd.Parse(os.Args[2:])
		id, err := strconv.Atoi(delCmd.Arg(0))
		if err != nil {
			log.Fatal("Неправильно указан id задачи")
		}
		if err := taskService.Delete(id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Задача %d удалена!\n", id)

	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
