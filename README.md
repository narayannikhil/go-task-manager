# CLI Task Manager

A command-line task management tool built with Go. Supports full CRUD operations with persistent JSON-based storage — no database or external dependencies required.

## Features

- Add tasks
- List all tasks with status
- Mark tasks as complete
- Delete tasks
- Data persists across sessions via a local `tasks.json` file

## Tech Stack

- **Language:** Go
- **Storage:** JSON file (no external DB)
- **Concepts used:** Structs, slices, file I/O, JSON encoding/decoding, CLI argument parsing

## Getting Started

### Prerequisites

- Go 1.18 or higher installed — [Download Go](https://golang.org/dl/)

### Installation

```bash
# Clone the repository
git clone https://github.com/narayannikhil/go-task-manager.git

# Navigate into the project
cd go-task-manager

# Initialize the module (if not already done)
go mod tidy
```

## Usage

```bash
# Add a new task
go run main.go add "Your task title"

# List all tasks
go run main.go list

# Mark a task as complete
go run main.go done <id>

# Delete a task
go run main.go delete <id>

# Show help
go run main.go help
```

## Example

```bash
$ go run main.go add "Learn Go structs"
✅ Task added: [1] Learn Go structs

$ go run main.go add "Push project to GitHub"
✅ Task added: [2] Push project to GitHub

$ go run main.go list

📋 Your Tasks:
-----------------------------
[ ] [1] Learn Go structs  (2026-03-24 11:00:00)
[ ] [2] Push project to GitHub  (2026-03-24 11:01:00)
-----------------------------

$ go run main.go done 1
✅ Task [1] marked as done: Learn Go structs

$ go run main.go list

📋 Your Tasks:
-----------------------------
[✓] [1] Learn Go structs  (2026-03-24 11:00:00)
[ ] [2] Push project to GitHub  (2026-03-24 11:01:00)
-----------------------------

$ go run main.go delete 2
🗑️  Deleted task [2]: Push project to GitHub
```

## Project Structure

```
go-task-manager/
├── main.go        # All application logic
├── go.mod         # Go module file
├── tasks.json     # Auto-generated on first task added
└── README.md      # Project documentation
```

## What I Learned

- Defining and using **structs** in Go
- Working with **slices** for in-memory data management
- **JSON encoding/decoding** for persistent file storage
- Parsing **CLI arguments** using `os.Args`
- File **read/write operations** in Go
- Basic **error handling** patterns

## Author

**Nikhil Pareek**
- GitHub: [@narayannikhil](https://github.com/narayannikhil)
- LinkedIn: [nikhilpareeknp](https://www.linkedin.com/in/nikhilpareeknp/)
