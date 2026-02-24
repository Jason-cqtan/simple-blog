# simple-blog

A simple blog system built with Go, Gin framework, GORM, and JWT authentication.

## Features

- **User Module** - Registration, login, logout, profile management
- **Blog Posts** - Create, edit, delete, list, detail, categories, tags
- **Comment System** - Add comments, list comments per post
- **Access Control** - JWT authentication, author permission control

## Tech Stack

| Component | Technology |
|-----------|-----------|
| Framework | [Gin](https://github.com/gin-gonic/gin) v1.9 |
| ORM | [GORM](https://gorm.io) v1.25 |
| Database | MySQL / PostgreSQL |
| Auth | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) v5 |
| Templates | Go `html/template` |

## Project Structure

```
simple-blog/
├── main.go                 # Single entry point
├── config/
│   └── config.go           # Environment-driven configuration
├── models/
│   ├── user.go             # User model (bcrypt password)
│   ├── post.go             # Post model
│   └── comment.go          # Comment model
├── routes/
│   └── routes.go           # Route definitions
├── handlers/
│   ├── user_handler.go     # User controller
│   ├── post_handler.go     # Post controller
│   └── comment_handler.go  # Comment controller
├── middleware/
│   └── auth.go             # JWT auth middleware
├── views/                  # HTML templates (html/template)
│   ├── layouts/base.html
│   ├── home.html
│   ├── posts/
│   ├── users/
│   └── comments/
├── static/                 # CSS, JS, images
├── database/
│   └── connection.go       # DB init + AutoMigrate
└── utils/
    ├── jwt.go              # JWT helpers
    └── validators.go       # Validation helpers
```

## Quick Start

### Prerequisites

- Go 1.21+
- MySQL or PostgreSQL

### Configuration

Copy the example environment variables and adjust as needed:

```bash
export DB_DRIVER=mysql          # mysql or postgres
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=yourpassword
export DB_NAME=simple_blog
export JWT_SECRET=your-secret-key
export SERVER_PORT=8080
export SECURE_COOKIE=false      # set true in production (HTTPS)
```

### Run

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The server starts on `http://localhost:8080` by default.

## API / Routes

| Method | Path | Description | Auth |
|--------|------|-------------|------|
| GET | `/` | Home page (recent posts) | No |
| GET | `/posts` | Post list | No |
| GET | `/posts/:id` | Post detail + comments | No |
| GET | `/register` | Register form | No |
| POST | `/register` | Submit registration | No |
| GET | `/login` | Login form | No |
| POST | `/login` | Submit login | No |
| POST | `/logout` | Logout | No |
| GET | `/posts/new` | Create post form | ✅ |
| POST | `/posts` | Submit new post | ✅ |
| GET | `/posts/:id/edit` | Edit post form | ✅ |
| PUT | `/posts/:id` | Submit post update | ✅ |
| DELETE | `/posts/:id` | Delete post | ✅ |
| POST | `/posts/:id/comments` | Add comment | ✅ |
| GET | `/profile` | User profile | ✅ |

## License

MIT
