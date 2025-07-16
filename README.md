# Task Manager API

A simple Task Management REST API built with Go and the Gin Framework.

## 🚀 Features

- Create, Read, Update, Delete (CRUD) tasks
- In-memory storage (no database yet)
- Structured by MVC architecture
- Tested with Postman and fully documented

🛠️ Tech Stack

- **Backend:** Go (Gin)
- **Tools:** Postman

## 📂 Folder Structure



task_manager/
├── main.go
├── controllers/
│   └── task\_controller.go
├── models/
│   └── task.go
├── data/
│   └── task\_service.go
├── router/
│   └── router.go
├── docs/
│   └── api\_documentation.md
└── go.mod

````

## 🧪 Running Locally

1. Clone the project  
```bash
git clone https://github.com/yourusername/go-task-manager.git
cd go-task-manager
````

2. Run the app

```bash
go run main.go
```

3. Server runs at:

```
http://localhost:8080
```

## 📖 API Documentation

Full API reference is available in [`docs/api_documentation.md`](./docs/api_documentation.md)

---

## ✅ To Do

* Add persistent database support
* User authentication
* Swagger/OpenAPI integration

---

## 📩 Contact

If you want to reach out, feel free to DM me [@surafelbkassa](https://t.me/surafelbkassa)
