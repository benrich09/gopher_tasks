main.go: Full working CLI todo app (fixed for clean build)
go.mod: Module initialized
README.md: Detailed instructions + extensions
.gitignore: Keeps repo clean


**✅ Files created successfully!**

You now have:
- `go.mod` — Go module file
- `tasks.json` — Initial empty data file

---

### 1. **go.mod** content

```go
module GopherTasks

go 1.23
```

This declares your project as a Go module named `GopherTasks`. It allows you to run `go run`, `go build`, `go test`, etc., cleanly and makes it easier to add dependencies later.

---

### 2. **How the Todo App Works (Detailed Explanation)**

Let me break down the entire program step by step.

#### **Core Data Structures**

```go
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
}
```

- `Task` represents one todo item.
- The backticks (`` `json:"..."` ``) are **JSON tags** — they tell Go how to map struct fields to JSON keys when saving/loading.

```go
type TodoList struct {
    Tasks  []Task `json:"tasks"`
    nextID int    // Not exported (lowercase) → not saved in JSON
}
```

- `TodoList` holds all tasks and manages the next available ID.

---

#### **Key Functions**

**`NewTodoList(filename string)`**
1. Creates a new `TodoList` with empty tasks and `nextID = 1`.
2. Tries to read `tasks.json`.
3. If file doesn't exist → starts fresh (no error).
4. If file exists → loads JSON into the struct using `json.Unmarshal`.
5. Scans all tasks to find the highest ID and sets `nextID` accordingly (so IDs don't collide after restart).

**`Save(filename string)`**
- Converts the entire `TodoList` to nicely formatted JSON (`json.MarshalIndent`).
- Writes it to `tasks.json`.

---

#### **CRUD Operations**

| Command   | Function       | What it does |
|-----------|----------------|--------------|
| `add`     | `Add()`        | Creates new `Task`, appends to slice, increments `nextID` |
| `list`    | `List()`       | Returns all tasks for display |
| `done`    | `Complete()`   | Finds task by ID and sets `Completed = true` |
| `remove`  | `Remove()`     | Finds task by ID and removes it from slice using `append` trick |

---

#### **How `main()` Works**

1. Loads the todo list from `tasks.json`.
2. Reads command-line arguments (`os.Args`).
3. Switches on the first argument (`add`, `list`, `done`, `remove`).
4. Performs the requested action.
5. **Always saves** the list at the end (even if nothing changed).

---

### 3. **How to Use It Now**

Make sure you're in the project folder and run:

```bash
# Add tasks
go run main.go add "Learn Go modules"
go run main.go add "Build a web API"

# List tasks
go run main.go list

# Mark as done
go run main.go done 1

# Remove a task
go run main.go remove 2

# See current tasks again
go run main.go list
```

---

### 4. **What happens under the hood when you run commands**

**Example: `go run main.go add "Buy milk"`**

1. `NewTodoList()` loads `tasks.json`.
2. `Add()` creates a new `Task` with next ID.
3. `Save()` writes updated JSON to disk.
4. You see: `Task added: Buy milk`

**Example: `go run main.go list`**

- Loads data
- Prints formatted table
- No save needed (read-only)

---

### 5. **Project Structure (Current)**

```
GopherTasks/
├── main.go
├── go.mod
├── tasks.json
└── README.md (you can create this)
```

Would you like me to also create a nice **README.md** with installation, usage examples, and future improvement ideas?

Also let me know if you want:
- Proper package structure (separate files)
- Unit tests
- CLI flags using `flag` package
- Or turn this into a web API (Gin/Fiber/echo)
# gopher_tasks
