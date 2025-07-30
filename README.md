Here’s your **updated `README.md`** reflecting your current progress, updated structure, CI test status, and tech improvements:

---

```markdown
# 🗂️ Task Manager API

A clean, modular, and test-driven Task Management REST API built with **Go** using the **Gin Framework**, following Clean Architecture principles.

---

## 🚀 Features

- ✅ Full CRUD operations for tasks
- ✅ Modular Clean Architecture (Domain, Usecases, Repositories, Delivery, Infrastructure)
- ✅ MongoDB integration (using the official MongoDB Go driver)
- ✅ Unit testing with `testify` and GitHub Actions CI
- ✅ Code coverage report generation
- 🛡️ Easy-to-extend for user auth, logging, middlewares, etc.

---

## 🧰 Tech Stack

- **Language:** Go 1.24+
- **Framework:** Gin
- **Database:** MongoDB
- **Testing:** `testify`, `mock`
- **CI/CD:** GitHub Actions (automated test + coverage pipeline)

---

## 📁 Folder Structure

```

task\_manager/
├── cmd/                # Entry point (main.go)
├── Domain/             # Entities and core business logic
├── Usecases/           # Application-specific use cases
├── Repositories/       # Repository interfaces
├── Infrastructure/     # External services (MongoDB)
├── Delivery/           # HTTP Handlers (Gin)
├── Tests/              # Unit tests
├── coverage.out        # Code coverage output
├── go.mod / go.sum     # Dependencies
└── .github/workflows/  # GitHub Actions CI

````

---

## 🧪 Running Locally

```bash
git clone https://github.com/surafelbkassa/go-task-manager.git
cd go-task-manager

# Run the app
go run cmd/main.go
````

Server will run at:
📍 `http://localhost:8080`

---

## 🔬 Testing

```bash
go test ./... -cover
```

### ✅ CI Status

![Go Test CI](https://github.com/surafelbkassa/go-task-manager/actions/workflows/go-test.yml/badge.svg)

> Automatically runs tests and generates code coverage reports on each push to `main`/`master`.

---

## 📖 API Documentation

> Documentation will be generated using Swagger in the next version.

---

## 🗓️ Roadmap / TODO

* [x] Add MongoDB persistent storage
* [x] Set up unit testing with mocking
* [x] Add GitHub Actions CI + code coverage
* [ ] Implement user authentication
* [ ] Add Swagger / OpenAPI spec
* [ ] Dockerize the app
* [ ] Deploy to Render/Vercel/Fly.io

---

## 📩 Contact

DM me on Telegram: [@surafelbkassa](https://t.me/surafelbkassa)
GitHub: [github.com/surafelbkassa](https://github.com/surafelbkassa)

---

Built with 💻 in Go.
