// main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Task represents a single todo item
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// TodoList manages a collection of tasks
type TodoList struct {
	Tasks  []Task `json:"tasks"`
	nextID int
}

// NewTodoList creates a new todo list, loading from file if it exists
func NewTodoList(filename string) (*TodoList, error) {
	tl := &TodoList{
		Tasks:  []Task{},
		nextID: 1,
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tl, nil // File doesn't exist yet - start fresh
		}
		return nil, err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &tl)
		if err != nil {
			return nil, err
		}
		// Calculate next ID
		for _, t := range tl.Tasks {
			if t.ID >= tl.nextID {
				tl.nextID = t.ID + 1
			}
		}
	}
	return tl, nil
}

// Save persists the todo list to a JSON file
func (tl *TodoList) Save(filename string) error {
	data, err := json.MarshalIndent(tl, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Add adds a new task
func (tl *TodoList) Add(description string) {
	task := Task{
		ID:          tl.nextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.nextID++
}

// List returns all tasks (for display)
func (tl *TodoList) List() []Task {
	return tl.Tasks
}

// Complete marks a task as done
func (tl *TodoList) Complete(id int) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

// Remove deletes a task
func (tl *TodoList) Remove(id int) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return true
		}
	}
	return false
}
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" go run main.go add \"Your task description\"")
	fmt.Println(" go run main.go list")
	fmt.Println(" go run main.go done <id>")
	fmt.Println(" go run main.go remove <id>")
	fmt.Println("\nData is saved automatically in tasks.json")
}
func main() {
	const filename = "tasks.json"
	tl, err := NewTodoList(filename)
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			printUsage()
			return
		}
		description := os.Args[2]
		tl.Add(description)
		fmt.Printf("Task added: %s\n", description)
	case "list":
		tasks := tl.List()
		if len(tasks) == 0 {
			fmt.Println("No tasks yet. Add one with: add \"Buy milk\"")
			return
		}
		fmt.Println("ID | Status | Description")
		fmt.Println("---|--------|------------")
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "✓"
			}
			fmt.Printf("%2d | %s | %s\n", t.ID, status, t.Description)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			printUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			return
		}
		if tl.Complete(id) {
			fmt.Printf("Task %d marked as completed!\n", id)
		} else {
			fmt.Printf("Task %d not found\n", id)
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide task ID")
			printUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			return
		}
		if tl.Remove(id) {
			fmt.Printf("Task %d removed\n", id)
		} else {
			fmt.Printf("Task %d not found\n", id)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
	// Always save after changes
	if err := tl.Save(filename); err != nil {
		fmt.Printf("Warning: Could not save tasks: %v\n", err)
	}
}
