# Golang Authentication Service

Microservice for user authentication and authorization built with Go.

## Features

- User registration and login
- JWT-based authentication
- Token refresh mechanism
- Password hashing with bcrypt
- Protected routes middleware
- Input validation

## Tech Stack

- **Go** - Programming language
- **PostgreSQL** - Database (or specify your DB)
- **JWT** - JSON Web Tokens for authentication
- **bcrypt** - Password hashing

## Installation

```bash
# Clone repository
git clone https://github.com/Pepega4a/Golang-auth_service.git
cd Golang-auth_service

# Install dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# Run the service
go run main.go
```

## Environment Variables

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_service
JWT_SECRET=your_secret_key
PORT=8080
```

## API Endpoints

### Authentication

**Register User**
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

**Login**
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Refresh Token**
```http
POST /api/auth/refresh
Authorization: Bearer {refresh_token}
```

### Protected Routes

**Get User Profile**
```http
GET /api/user/profile
Authorization: Bearer {access_token}
```

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── repository/
├── pkg/
│   └── utils/
├── go.mod
├── go.sum
└── README.md
```

## Security

- Passwords are hashed using bcrypt
- JWT tokens with expiration
- Environment variables for sensitive data
- Input validation and sanitization

## License

MIT

## Author

Vsevolod - [GitHub](https://github.com/Pepega4a)
