````markdown
# Users CRUD API 🧑‍💻

A RESTful API built in Go for user management with JWT authentication, role-based access control (admin/user), request logging to MongoDB, and Swagger documentation.

## ✨ Features

- User registration and login
- JWT authentication with token generation
- Middleware-protected routes
- Role-based access control (`admin` / `user`)
- Update/delete your own user account
- Admin-only role promotion endpoint
- Request logging to MongoDB (via Logrus)
- Swagger documentation with live interface

## 🧰 Middleware Overview

This project uses several key middlewares to enhance security, traceability, and reliability:

### 📘 `MiddlewareLogger()`
Logs all incoming requests using Logrus. Each request is logged both to the terminal and to MongoDB. Includes:
- HTTP method, path, status, latency
- User ID (if authenticated)
- Request body and headers (when necessary)

### 🛡 `MiddlewareRecovery()`
Gracefully handles panics and prevents server crashes. Automatically returns a 500 response when unexpected errors occur, logging the stack trace.

### 🌐 `MiddlewareCORS()`
Enables Cross-Origin Resource Sharing so the API can be consumed from web frontends on different origins (e.g. React apps running on localhost).

### 🚦 `MiddlewareRateLimit()`
Basic rate limiting to prevent abuse (e.g. too many login attempts). Helps reduce load and brute-force risk.

Each middleware is registered globally in `server.Router()`:

r.Use(MiddlewareLogger())
r.Use(MiddlewareRecovery())
r.Use(MiddlewareCORS())
r.Use(MiddlewareRateLimit())
````

---

## 🛠️ Tech Stack

* [Go 1.24+](https://go.dev/)
* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* PostgreSQL (for persistence)
* MongoDB (for structured logging)
* JWT (for authentication)
* Swaggo (Swagger auto-gen)
* Docker (for containerization)

---

## 📁 Project Structure

```
users-crud/
├── cmd/              # App entry point
├── internal/
│   ├── config/       # DB and Mongo config
│   ├── dto/          # Data transfer objects
│   ├── server/       # Middlewares and route setup
│   ├── handlers/     # HTTP request handlers
│   ├── models/       # GORM models
│   ├── repositories/ # Data access logic
│   ├── usecases/     # Business rules
│   ├── dto/          # Request/response structures
│   ├── logger/       # Logrus + Mongo integration
│   └── middleware/   # JWT, CORS, rate-limiters
├── docs/             # Swagger auto-generated files
├── go.mod
├── Dockerfile
└── .env
```

---

## 🚀 How to Run

### ✅ Requirements

* PostgreSQL running on `localhost:5432`
* MongoDB running on `localhost:27017`
* Go 1.24+ installed

---

### 📄 Environment Variables (`.env`)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=crud

MONGO_URI=mongodb://admin:admin@localhost:27017
```

---

### 🧪 Run Locally

```bash
go run ./cmd/main.go
```

---

### 🐳 Run with Docker

```bash
docker build -t users-api .
docker run --env-file .env -p 8080:8080 users-api
```

---

## 📄 Swagger API Documentation

Access Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

---

## 🔐 Authentication

* After login, you'll receive a JWT token
* Use `Authorization: Bearer <token>` in all protected requests

---

## 📬 Endpoints Overview

| Method | Endpoint                    | Description             | Auth    |
| ------ | --------------------------- | ----------------------- | ------- |
| POST   | `/api/register`             | Register a new user     | ❌       |
| POST   | `/api/login`                | Login and get JWT token | ❌       |
| GET    | `/api/me`                   | Get own user data       | ✅       |
| PUT    | `/api/users/:id`            | Update own user info    | ✅       |
| DELETE | `/api/users/:id`            | Delete own user account | ✅       |
| PUT    | `/api/admin/users/:id/role` | Promote/demote role     | ✅ admin |

---

## 🧾 Logging

All requests are logged to:

* Terminal (with Logrus)
* MongoDB (`logs` collection)
