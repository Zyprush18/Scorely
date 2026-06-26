# Scorely

A REST API for managing exams, student scores, and grading with role-based access control. Built with Go, GORM, JWT, and Redis.

## Features

- **Authentication & Authorization** — JWT access tokens with Redis-backed refresh token rotation. Automatic token renewal via middleware when the access token expires.
- **Role-Based Access Control** — Three roles: `admin`, `teacher`, `student`. Middleware enforces role restrictions per endpoint.
- **CRUD Management** — Full CRUD for roles, users, majors, levels, classes, subjects, teachers, students, exams, and exam questions.
- **Grading System** — Exam questions with multiple-choice options, correct-answer tracking, and answer scoring per student.
- **Multi-Database Support** — MySQL, PostgreSQL, SQLite, and GaussDB via GORM.
- **Input Validation** — Request validation with field-level error messages.
- **Structured Logging** — File-based error logging.

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.26 |
| HTTP | `net/http` with `http.ServeMux` (Go 1.22+ routing) |
| Database ORM | GORM (MySQL, PostgreSQL, SQLite, GaussDB) |
| Cache | Redis (refresh token storage) |
| Auth | JWT (HS256) with access + refresh tokens |
| Validation | `go-playground/validator` |
| Env Config | `joho/godotenv` |
| Testing | `testify`, `sqlmock` |

## Project Structure

```
Scorely/
├── cmd/
│   └── main.go                  # Entry point — loads env, connects DB & Redis, starts server
├── config/
│   ├── config.go                # Loads .env file
│   ├── jwt.go                   # (moved to helper/jwt.go)
│   └── redis.go                 # ConnectRedis() — creates *redis.Client
├── database/
│   └── connect.go               # DB connection & auto-migration
├── handlers/
│   ├── auth/                    # Login, Signup
│   ├── class/                   # Class CRUD
│   ├── exam/                    # Exam CRUD
│   ├── examquestion/            # Exam question CRUD
│   ├── level/                   # Level CRUD
│   ├── major/                   # Major CRUD
│   ├── role/                    # Role CRUD
│   ├── student/                 # Student CRUD
│   ├── subject/                 # Subject CRUD
│   ├── teacher/                 # Teacher CRUD
│   └── user/                    # User CRUD
├── helper/
│   ├── helper.go                # HTTP helpers, validation, pagination, password hashing
│   └── jwt.go                   # JWT generation, parsing, custom claims
├── middleware/
│   └── auth_middleware.go       # JWT auth + role check + refresh token rotation
├── models/
│   ├── entity/                  # GORM entity models
│   ├── request/                 # Request DTOs with validation tags
│   └── response/                # Response DTOs
├── repository/                  # Data access layer (GORM)
│   ├── repoauth/
│   ├── repoclass/
│   ├── repoexamquestions/
│   ├── repoexams/
│   ├── repolevel/
│   ├── repomajor/
│   ├── reporole/
│   ├── repostudent/
│   ├── reposubject/
│   ├── repoteacher/
│   └── repouser/
├── routes/
│   └── web.go                   # All route definitions
├── service/                     # Business logic layer
│   ├── classservice/
│   ├── majorservice/
│   ├── serviceauth/
│   ├── serviceexam/
│   ├── serviceexamquest/
│   ├── servicelevel/
│   ├── servicerole/
│   ├── servicestudent/
│   ├── serviceteacher/
│   ├── subjectservice/
│   └── userservice/
├── .env                         # Environment variables
├── go.mod
├── go.sum
└── README.md
```

## Environment Variables

Create a `.env` file in the project root:

```env
# Database
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=scorely
DB_USERNAME=root
DB_PASSWORD=

# JWT
JWT_SECRET_KEY=your-jwt-secret
REFRESH_SECRET_KEY=your-refresh-secret

# Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

## Getting Started

### Prerequisites

- Go 1.26+
- MySQL (or PostgreSQL, SQLite, GaussDB)
- Redis

### Run

```bash
# Clone the repository
git clone <repo-url>
cd Scorely

# Configure environment
cp .env.example .env
# Edit .env with your database and Redis credentials

# Install dependencies
go mod tidy

# Run
go run cmd/main.go
```

The server starts on `http://localhost:8000`.

## API Endpoints

### Authentication (public)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/login` | Login — returns JWT access token |
| POST | `/api/register` | Register a new user |

### Roles (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/role` | List roles (paginated) |
| POST | `/api/role/add` | Create role |
| GET | `/api/role/{id}` | Get role by ID |
| PUT | `/api/role/{id}/update` | Update role |
| DELETE | `/api/role/{id}/delete` | Delete role |

### Users (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/user` | List users (paginated) |
| POST | `/api/user/add` | Create user |
| GET | `/api/user/{id}` | Get user by ID |
| PUT | `/api/user/{id}/update` | Update user |
| DELETE | `/api/user/{id}/delete` | Delete user |

### Majors (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/major` | List majors (paginated) |
| POST | `/api/major/add` | Create major |
| GET | `/api/major/{id}` | Get major by ID |
| PUT | `/api/major/{id}/update` | Update major |
| DELETE | `/api/major/{id}/delete` | Delete major |

### Levels (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/level` | List levels (paginated) |
| POST | `/api/level/add` | Create level |
| GET | `/api/level/{id}` | Get level by ID |
| PUT | `/api/level/{id}/update` | Update level |
| DELETE | `/api/level/{id}/delete` | Delete level |

### Classes (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/class` | List classes (paginated) |
| POST | `/api/class/add` | Create class |
| GET | `/api/class/{id}` | Get class by ID |
| PUT | `/api/class/{id}/update` | Update class |
| DELETE | `/api/class/{id}/delete` | Delete class |

### Students (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/student` | List students (paginated) |
| POST | `/api/student/add` | Create student |
| GET | `/api/student/{id}` | Get student by ID |
| PUT | `/api/student/{id}/update` | Update student |
| DELETE | `/api/student/{id}/delete` | Delete student |

### Teachers (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/teacher` | List teachers (paginated) |
| POST | `/api/teacher/add` | Create teacher |
| GET | `/api/teacher/{id}` | Get teacher by ID |
| PUT | `/api/teacher/{id}/update` | Update teacher |
| DELETE | `/api/teacher/{id}/delete` | Delete teacher |

### Subjects (admin)

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/subject` | List subjects (paginated) |
| POST | `/api/subject/add` | Create subject |
| GET | `/api/subject/{id}` | Get subject by ID |
| PUT | `/api/subject/{id}/update` | Update subject |
| DELETE | `/api/subject/{id}/delete` | Delete subject |

### Exams (admin, teacher)

| Method | Path | Roles | Description |
|--------|------|-------|-------------|
| GET | `/api/exam` | admin | List all exams (paginated) |
| GET | `/api/teacher/exam` | teacher | List exams by teacher (from JWT) |
| POST | `/api/exam/{subject_id}/add` | admin, teacher | Create exam for a subject |
| GET | `/api/exam/{id}` | admin, teacher | Get exam by ID |
| PUT | `/api/exam/{id}/update` | admin, teacher | Update exam |
| DELETE | `/api/exam/{id}/delete` | admin, teacher | Delete exam |

### Exam Questions (admin, teacher)

| Method | Path | Roles | Description |
|--------|------|-------|-------------|
| GET | `/api/exam/{id_exam}/examquestion` | admin, teacher | List questions for an exam |
| POST | `/api/exam/{id_exam}/examquestion/add` | admin, teacher | Add question to an exam |
| GET | `/api/exam/{id_exam}/examquestion/{id}` | admin, teacher | Get question by ID |

## Authentication Flow

1. **Login** — `POST /api/login` with email/password returns a JWT access token (24h expiry). A refresh token (7d expiry) is also generated and stored in Redis as `refresh_token:{user_id}`.

2. **Authenticated Requests** — Include the access token in the `Authorization: Bearer <token>` header. The middleware validates the token signature, checks expiry, and verifies the user's role against the endpoint's allowed roles.

3. **Token Refresh** — When the access token is expired but the signature is valid, the middleware automatically:
   - Extracts the user ID from the expired token's claims
   - Retrieves the refresh token from Redis (`refresh_token:{user_id}`)
   - Validates the refresh token's signature and expiry
   - Generates a new access token and returns it in the `X-New-Token` response header
   - Continues processing the original request

4. **Logout** — Delete the refresh token from Redis to invalidate the session.

## Response Format

All endpoints return JSON with a consistent envelope:

```json
{
  "message": "Success",
  "data": {},
  "token": "",
  "error": "",
  "field": "",
  "pagination": {
    "page": 1,
    "total_data": 100,
    "perpage": 10,
    "total_page": 10,
    "prev": "/role?page=0",
    "next": "/role?page=2"
  }
}
```
