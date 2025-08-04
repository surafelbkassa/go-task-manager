# Task Manager REST API (Go + Gin)

A fully functional Task Management API built with Go and Gin, following modern backend development practices and clean architecture principles.

This project was developed in progressive stages:
- ✅ CRUD API using Gin and in-memory storage
- ✅ MongoDB integration for persistence
- ✅ JWT-based authentication and role-based access control
- ✅ Refactored into clean architecture for maintainability
- ✅ Unit tested with `testify`

---

## 📌 Features

- **User Registration & Login**
  - Secure password hashing
  - JWT token generation and validation

- **Task Management**
  - Create, update, delete, and retrieve tasks
  - Each task has: `title`, `description`, `due date`, `status`
  - Ownership enforcement with user ID linkage

- **Authentication & Authorization**
  - Middleware-protected routes
  - Role-based access (e.g., Admin-only routes)

- **Clean Architecture**
  - Domain-driven structure
  - Interfaces for repositories and services
  - Separation of concerns: Delivery, Usecase, Repository, Infrastructure

- **MongoDB Integration**
  - Replaces in-memory storage
  - Tasks and users are persisted using the official Mongo Go driver

- **Testing**
  - Unit tests for services and use cases using `testify`
  - Mocked dependencies for isolated testing

---

## 🧱 Project Structure

```bash
task-manager/
├── Delivery/
├── Domain/
├── Infrastructure/
├── Repositories/
├── Usecases/
├── models/
├── docs/
│   └── api_documentation.md
└── go.mod

````

---

## 🔌 Tech Stack

* **Backend**: Go (Gin)
* **Database**: MongoDB (using mongo-go-driver)
* **Auth**: JWT
* **Testing**: Testify
* **Docs**: Postman

---

## 🚀 How to Run

1. Clone the repo:

   ```bash
   git clone https://github.com/surafelbkassa/task-manager-api.git
   cd task-manager-api
   ```

2. Set up MongoDB (local or cloud) and update the `.env` or connection string.

3. Run the server:

   ```bash
   go run Delivery/main.go
   ```

---

## 🔐 Authentication Flow

1. Register: `POST /register`
2. Login: `POST /login`
3. Receive JWT Token
4. Access protected routes with `Authorization: Bearer <token>`

---

## 🧪 Testing

Run unit tests:

```bash
go test ./... -v
```

Mocking is done using `testify/mock` to isolate logic.

---

## 📄 API Documentation

Full documentation (with examples and test cases) available in the [Postman Docs](https://documenter.getpostman.com/view/38360301/2sB3BBpBEc) and in `docs/api_documentation.md`.

---

## 💼 Use Case

This backend can serve as the foundation for:

* A personal productivity app
* A SaaS platform for teams
* A microservice inside a larger system

It’s structured to be extensible — features like notifications, task sharing, and analytics can be added with minimal friction.

---

## 👨‍💻 Author

Built by Surafel — backend-focused developer committed to writing clean, scalable, and secure APIs.
Connect via [LinkedIn](www.linkedin.com/in/surafelbkassa) | [Email](surafelbkassa3@gmail.com) | [Upwork Profile](https://www.upwork.com/freelancers/~01c0f840cec272f38a)

---

## 📌 Note

This project was built as part of a deep-dive backend training, and has been incrementally upgraded to meet real-world production standards. Every piece was designed for practical usage and long-term maintainability.

```
