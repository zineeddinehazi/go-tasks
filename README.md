# 🚀 Go-Tasks: Task Management API

Welcome to **Go-Tasks** ✨, a lightweight RESTful API built with Go, Gorilla Mux, and SQLite for managing your tasks efficiently. Perfect for learning Go backend development, practicing REST design, and building portfolio-ready projects 🛠️.

---

## 🌟 Features

- **Persistent SQLite storage** 🗄️: Tasks saved in `tasks.db` file (auto-created)
- **Full CRUD operations** 📤: Create, read, update, toggle status, delete tasks
- **Repository pattern** 🧱: Clean separation of database logic in `TaskRepository`
- **Structured project layout** 📂: Organized packages for handlers, router, repository
- **RESTful endpoints** 🌐: Predictable URLs and proper HTTP status codes
- **JSON I/O** 📦: Seamless frontend integration
- **Toggle completion** ✅: Single request flips `isdone` boolean in database

---

## 📡 Endpoints

| Method   | Endpoint     | Description                  |
| -------- | ------------ | ---------------------------- |
| `GET`    | `/list`      | Get all tasks 📋             |
| `GET`    | `/list/{id}` | Get a specific task by ID 🔍 |
| `POST`   | `/list`      | Add a new task ➕            |
| `PATCH`  | `/list/{id}` | Update a task's status ✅    |
| `DELETE` | `/list/{id}` | Delete a task by ID ❌       |

---

## 🚀 Getting Started

### Prerequisites
- Go 1.20+ 
- `go get github.com/gorilla/mux`
- `go get github.com/mattn/go-sqlite3`

### Installation

1. Clone the repository:

```bash
git clone https://github.com/zineeddinehazi/go-tasks.git
```

2. Run the server:

```bash
cd go-tasks
go run cd/main.go
```

3. Access the API at `http://localhost:8080` 🌍.

---

## 💡 Usage Examples

### Add a Task ➕

```bash
curl -X POST http://localhost:8080/list -H "Content-Type: application/json" -d '{"content": "Write documentation", "isdone": false}'
```

### Update a Task ✅

```bash
curl -X PATCH http://localhost:8080/list/5 -H "Content-Type: application/json" -d '{}'
```

### Delete a Task ❌

```bash
curl -X DELETE http://localhost:8080/list/5
```

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for new features, bug fixes, or improvements 🛠️.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details 📄.

---

## 🙏 Acknowledgments

Built with Go 🐹, Gorilla Mux 🧩, and SQLite 🗄️. Inspired by the need for simple, reliable, and beautiful APIs 💡.

---

Thank you for using Go-Tasks! If you have any questions or suggestions, feel free to reach out 📩. Happy coding! 💻✨
