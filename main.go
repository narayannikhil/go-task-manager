package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
}

const dataFile = "tasks.json"

func loadTasks() []Task {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(dataFile, data, 0644)
}

func addTask(title string) {
	tasks := loadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	task := Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("✅ Task added: [%d] %s\n", task.ID, task.Title)
}

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found. Add one with: go run main.go add \"your task\"")
		return
	}
	fmt.Println("\n📋 Your Tasks:")
	fmt.Println("-----------------------------")
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[✓]"
		}
		fmt.Printf("%s [%d] %s  (%s)\n", status, t.ID, t.Title, t.CreatedAt)
	}
	fmt.Println("-----------------------------")
}

func completeTask(id int) {
	tasks := loadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			saveTasks(tasks)
			fmt.Printf("✅ Task [%d] marked as done: %s\n", id, t.Title)
			return
		}
	}
	fmt.Printf("❌ Task [%d] not found\n", id)
}

func deleteTask(id int) {
	tasks := loadTasks()
	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			fmt.Printf("🗑️  Deleted task [%d]: %s\n", id, t.Title)
		} else {
			newTasks = append(newTasks, t)
		}
	}
	if !found {
		fmt.Printf("❌ Task [%d] not found\n", id)
		return
	}
	saveTasks(newTasks)
}

func printHelp() {
	fmt.Println(`
CLI Task Manager — Built with Go

Usage:
  go run main.go add "task title"     Add a new task
  go run main.go list                 List all tasks
  go run main.go done <id>            Mark a task as complete
  go run main.go delete <id>          Delete a task
  go run main.go help                 Show this help message
`)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: go run main.go add \"task title\"")
			return
		}
		addTask(args[1])
	case "list":
		listTasks()
	case "done":
		if len(args) < 2 {
			fmt.Println("Usage: go run main.go done <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("❌ Please provide a valid task ID (number)")
			return
		}
		completeTask(id)
	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: go run main.go delete <id>")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("❌ Please provide a valid task ID (number)")
			return
		}
		deleteTask(id)
	case "help":
		printHelp()
	default:
		fmt.Printf("❌ Unknown command: %s\n", args[0])
		printHelp()
	}
}




