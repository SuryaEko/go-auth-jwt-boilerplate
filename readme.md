# Go Auth JWT Boilerplate

A clean and modular boilerplate for building RESTful APIs in Go using Gin, GORM, and JWT authentication.

## Features
- User registration & login with JWT authentication (secure password hashing with bcrypt)
- Role-based access (admin, user)
- JWT token generation and validation
- PostgreSQL/SQLite/MySQL support via GORM
- Modular and scalable folder structure (models, services, controllers, routes, middleware, utils)
- Environment variable configuration via `.env`
- Dependency Injection ready (service container)
- Auto-migration for database tables
- Example for profile update and password change

## Project Structure
```
go-auth-jwt-boilerplate/
├── main.go                # Application entry point
├── go.mod / go.sum        # Go modules
├── .env.example           # Environment variables (example)
│
├── database/              # DB connection & auto migration
├── models/                # GORM models (User, etc)
├── dto/                   # Data Transfer Objects (input/output struct)
├── services/              # Business logic (user, auth, profile, etc)
├── controllers/           # HTTP handlers (register, login, profile, etc)
├── routes/                # Route registration
├── middleware/            # JWT & custom middleware
├── utils/                 # Helper functions (JWT, bcrypt, etc)
```

## Getting Started

### 1. Clone the repository
```sh
git clone https://github.com/SuryaEko/go-auth-jwt-boilerplate.git
cd go-auth-jwt-boilerplate
```

### 2. Setup environment variables
Copy `.env.example` to `.env` and fill in your configuration:
```sh
cp .env.example .env
```

Edit `.env` as needed:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
DB_PORT=5432
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRATION=24h
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
- `PUT /profile` — Update user profile (protected)
- `PUT /profile/password` — Change user password (protected)
- `POST /users` — Create a new user (protected)
- `GET /users/:id` — Get user by ID (protected)
- `PUT /users/:id` — Update user by ID (protected)
- `PUT /users/:id/password` — Update user password by ID (protected)

## Tools & Libraries
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://gorm.io/) — ORM for Go
- [JWT](https://github.com/golang-jwt/jwt) — JSON Web Token
- [godotenv](https://github.com/joho/godotenv) — Load env file
- [crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing

## Articles
- [Gorm Pagination](https://dev.to/rafaelgfirmino/pagination-using-gorm-scopes-3k5f) - How to implement pagination using GORM scopes

## License
MIT