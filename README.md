# 🏦 Payment System API (GoLang)

A scalable, JWT-secured user authentication service with rate-limiting using Go, PostgreSQL, and Gorilla Mux. Built following a clean, modular, production-ready architecture.

---

## 📁 Project Structure

```
payment-system/
├── cmd/
│   └── api/
│       └── main.go            # App entry point
├── internal/
│   ├── common/
│   │   └── db/                # Database connection logic
│   ├── config/                # Loads .env and config values
│   ├── middleware/            # JWT + Rate limiting
│   └── user/                  # User domain: handler, service, repository, model
├── .env                       # Environment config
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## 💠 Tech Stack

- **Go 1.20+**
- **PostgreSQL**
- **Gorilla Mux** (routing)
- **bcrypt** (password hashing)
- **JWT** (auth)
- **godotenv** (env file support)
- **Rate Limiter** (custom middleware)

---

## ⚙️ Setup Instructions

### 1. Clone the project

```bash
git clone https://github.com/your-username/payment-system.git
cd payment-system
```

### 2. Create `.env` file

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=paymentdb
JWT_SECRET=your-secret-key
```

### 3. Initialize Go Modules

```bash
go mod tidy
```

### 4. Run the server

```bash
go run cmd/api/main.go
```

---

## 🥪 API Endpoints

| Method | Endpoint    | Auth Required | Description          |
| ------ | ----------- | ------------- | -------------------- |
| POST   | `/register` | ❌             | Register new user    |
| POST   | `/login`    | ❌             | Login user (get JWT) |
| GET    | `/users`    | ✅             | Get all users        |

---

## 🔐 Authentication

All routes **except **``** and **`` require a valid JWT.

### Login Response

```json
{
  "token": "your.jwt.token"
}
```

### Use token in header

```
Authorization: Bearer your.jwt.token
```

---

## 🧱 Rate Limiting

Custom middleware limits the number of requests per client (e.g., 100 req/min). Can be customized inside `internal/middleware/rate_limit.go`.

---

## 💃 Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL
);
```

> You can use a migration tool like `golang-migrate`, or run SQL manually.

---

## 📀 .gitignore

Make sure to exclude sensitive files:

```
.env
*.exe
*.out
*.log
vendor/
```

---

## 📌 TODO (next steps)

-

---

## 🧑‍💼 Author

Built with ❤️ by [Your Name]

---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).

