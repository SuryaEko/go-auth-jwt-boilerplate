# Go Auth JWT Boilerplate

A clean and modular boilerplate for building RESTful APIs in Go using Gin, GORM, and JWT authentication.

## Features
- User registration & login with JWT authentication
- Password hashing (bcrypt)
- PostgreSQL/SQLite/MySQL support via GORM
- Modular structure: models, services, controllers, routes, middleware
- Environment variable support via `.env`
- Ready for Dependency Injection (DI)

## Project Structure
```
go-auth-jwt-boilerplate/
├── main.go                # Entry point
├── go.mod / go.sum        # Go modules
├── .env                   # Environment variables
│
├── database/              # DB connection & auto migration
├── models/                # GORM models
├── services/              # Business logic (user, auth, etc)
├── controllers/           # HTTP handlers
├── routes/                # Route registration
├── middleware/            # JWT middleware
├── utils/                 # Helper functions (JWT)
```

## Getting Started

### 1. Clone the repository
```sh
git clone https://github.com/yourusername/go-auth-jwt-boilerplate.git
cd go-auth-jwt-boilerplate
```

### 2. Setup environment variables
Create a `.env` file in the project root:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

### 3. Install dependencies
```sh
go mod tidy
```

### 4. Run the application
```sh
go run main.go
```

The app will run at `http://localhost:8080`

## API Endpoints (example)
- `POST /register` — Register a new user
- `POST /login` — Login and get JWT
- `GET /profile` — Get user profile (protected)

## Tools & Libraries
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://gorm.io/) — ORM for Go
- [JWT](https://github.com/golang-jwt/jwt) — JSON Web Token
- [godotenv](https://github.com/joho/godotenv) — Load env file
- [crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing