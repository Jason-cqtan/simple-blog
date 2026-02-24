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
├── main.go                 # Entry point (supports --seed flag)
├── .env.example            # Environment variable template
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
│   ├── connection.go       # DB init + AutoMigrate
│   ├── seeder.go           # Seed data (admin user, sample posts/comments)
│   └── migrations/
│       └── 001_initial_schema.sql  # Reference schema (PostgreSQL)
├── scripts/
│   └── init-db.sql         # PostgreSQL database/user bootstrap script
└── utils/
    ├── jwt.go              # JWT helpers
    └── validators.go       # Validation helpers
```

## Quick Start

### Prerequisites

- Go 1.21+
- MySQL or PostgreSQL

### Configuration

Copy the example environment file and adjust values:

```bash
cp .env.example .env
# Edit .env with your database credentials and secret key
set -a; source .env; set +a
```

Or export the variables manually:

```bash
export DB_DRIVER=postgres       # postgres or mysql
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=blog_user
export DB_PASSWORD=your_db_password
export DB_NAME=simple_blog
export JWT_SECRET=your-secret-key
export SERVER_PORT=8080
export SECURE_COOKIE=false      # set true in production (HTTPS)
```

### PostgreSQL Setup

1. Run the bootstrap script as a PostgreSQL superuser to create the database and
   application user:

   ```bash
   psql -U postgres -f scripts/init-db.sql
   ```

2. The application will create all tables automatically via GORM AutoMigrate on
   first startup. The reference schema is documented in
   `database/migrations/001_initial_schema.sql`.

### Run

```bash
# Install dependencies
go mod tidy

# Start the server (tables are created automatically on first run)
go run main.go

# Seed the database with an admin user and sample content, then exit
go run main.go --seed
```

After seeding, you can log in with:

| Field    | Value              |
|----------|--------------------|
| Username | `admin`            |
| Password | `Admin@123456`     |

> **Security notice:** Change the default admin password immediately after your first login,
> especially before any non-local or production deployment.

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
