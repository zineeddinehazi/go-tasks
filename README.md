# 🚀 Go-Tasks: Task Management API

Welcome to **Go-Tasks** ✨, a lightweight and powerful RESTful API built with Go and Gorilla Mux for managing your tasks efficiently. Whether you're learning Go, building a portfolio, or integrating with a frontend, this API provides a clean, reliable, and fully documented backend for your task management needs 🛠️.

---

## 🌟 Features

- **Full CRUD Operations** 📥📤: Create, read, update, and delete tasks with ease.
- **RESTful Design** 🌐: Follows REST conventions for predictable and scalable endpoints.
- **Error Handling** ⚠️: Detailed error messages and proper HTTP status codes for every scenario.
- **JSON Responses** 📦: All data is served in JSON format for seamless frontend integration.
- **Easy Setup** ⚡: Minimal dependencies and clear documentation.

---

## 📡 Endpoints

| Method   | Endpoint        | Description                      |
|----------|-----------------|----------------------------------|
| `GET`    | `/list`         | Get all tasks 📋                 |
| `GET`    | `/list/{id}`    | Get a specific task by ID 🔍     |
| `POST`   | `/list`         | Add a new task ➕                |
| `PATCH`  | `/list/{id}`    | Update a task's status ✅        |
| `DELETE` | `/list/{id}`    | Delete a task by ID ❌           |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher 🐹
- Gorilla Mux (`go get github.com/gorilla/mux`) 🧩

### Installation

1. Clone the repository:
```bash
git clone https://github.com/zineeddinehazi/go-tasks.git
```
2. Run the server:
```bash
cd go-tasks
go run main.go
```
3. Access the API at `http://localhost:8080` 🌍.

---

## 💡 Usage Examples

### Add a Task ➕
```bash
curl -X POST http://localhost:8080/list -H "Content-Type: application/json" -d '{"id": "5", "content": "Write documentation", "isdone": false}'
```
### Update a Task ✅
```bash
curl -X PATCH http://localhost:8080/list/5 -H "Content-Type: application/json" -d '{"isdone": true}'
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

Built with Go 🐹 and Gorilla Mux 🧩. Inspired by the need for simple, reliable, and beautiful APIs 💡.

---

Thank you for using Go-Tasks! If you have any questions or suggestions, feel free to reach out 📩. Happy coding! 💻✨

```
