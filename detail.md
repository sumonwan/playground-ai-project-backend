# Backend API Documentation

## Overview
This is a backend API built with Go Fiber framework that provides user authentication functionality with JWT tokens. The system includes user registration, login, and database management using GORM with SQLite.

## System Architecture

### Components
- **Go Fiber**: Web framework for handling HTTP requests
- **GORM**: ORM for database operations
- **SQLite**: Database for storing user data
- **JWT**: Token-based authentication
- **Bcrypt**: Password hashing
- **Swagger**: API documentation

## Sequence Diagrams

### User Registration Flow

```mermaid
sequenceDiagram
    participant Client
    participant API as Fiber API
    participant Auth as Auth Service
    participant DB as Database
    
    Client->>API: POST /auth/register
    Note over Client,API: {email, password}
    
    API->>API: Parse request body
    API->>DB: Check if user exists
    alt User already exists
        DB-->>API: User found
        API-->>Client: 409 Conflict
    else User not found
        DB-->>API: User not found
        API->>Auth: Hash password
        Auth-->>API: Hashed password
        API->>DB: Create new user
        DB-->>API: User created
        API->>Auth: Generate JWT token
        Auth-->>API: JWT token
        API-->>Client: 201 Created + Token + User
    end
```

### User Login Flow

```mermaid
sequenceDiagram
    participant Client
    participant API as Fiber API
    participant Auth as Auth Service
    participant DB as Database
    
    Client->>API: POST /auth/login
    Note over Client,API: {email, password}
    
    API->>API: Parse request body
    API->>DB: Find user by email
    alt User not found
        DB-->>API: No user found
        API-->>Client: 401 Unauthorized
    else User found
        DB-->>API: User data
        API->>Auth: Check password
        alt Invalid password
            Auth-->>API: Password mismatch
            API-->>Client: 401 Unauthorized
        else Valid password
            Auth-->>API: Password valid
            API->>Auth: Generate JWT token
            Auth-->>API: JWT token
            API-->>Client: 200 OK + Token + User
        end
    end
```

### API Request Flow

```mermaid
sequenceDiagram
    participant Client
    participant Middleware as CORS/Logger
    participant Router as Fiber Router
    participant Handler as Route Handler
    participant DB as Database
    
    Client->>Middleware: HTTP Request
    Middleware->>Router: Process request
    Router->>Handler: Route to handler
    alt Database operation needed
        Handler->>DB: Query/Update
        DB-->>Handler: Result
    end
    Handler-->>Router: Response
    Router-->>Middleware: Response
    Middleware-->>Client: HTTP Response
```

## Entity Relationship Diagram

```mermaid
erDiagram
    USER {
        uint id PK "Primary Key, Auto Increment"
        string email UK "Unique, Not Null"
        string password "Not Null, Hashed with bcrypt"
        datetime created_at "Auto Generated"
        datetime updated_at "Auto Generated"
        datetime deleted_at "Soft Delete, Nullable"
    }
    
    JWT_TOKEN {
        string token "JWT Token String"
        uint user_id FK "References USER.id"
        datetime expires_at "Token Expiration"
        datetime issued_at "Token Issue Time"
    }
    
    USER ||--o{ JWT_TOKEN : "generates"
```

## Data Models

### User Model
```go
type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Email    string `json:"email" gorm:"unique;not null"`
    Password string `json:"-" gorm:"not null"`
}
```

### Request/Response Models
```go
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}
```

## API Endpoints

### Authentication Routes

| Method | Endpoint | Description | Request Body | Response |
|--------|----------|-------------|--------------|----------|
| POST | `/auth/register` | Register new user | `RegisterRequest` | `AuthResponse` |
| POST | `/auth/login` | Login user | `LoginRequest` | `AuthResponse` |
| GET | `/` | Health check | None | `{message: "hello world"}` |
| GET | `/swagger/*` | API Documentation | None | Swagger UI |

## Security Features

1. **Password Hashing**: Uses bcrypt with default cost for secure password storage
2. **JWT Authentication**: Tokens expire in 72 hours
3. **Input Validation**: Email format and password minimum length validation
4. **CORS Protection**: Cross-origin request handling
5. **Request Logging**: All requests are logged for monitoring

## Database Schema

The system uses SQLite database with GORM for the following benefits:
- **Auto Migration**: Database schema is automatically created/updated
- **Soft Deletes**: Records are marked as deleted instead of being removed
- **Timestamps**: Automatic created_at and updated_at tracking
- **Relationships**: Easy foreign key management

## Technology Stack

- **Backend Framework**: Go Fiber v2
- **Database ORM**: GORM
- **Database**: SQLite
- **Authentication**: JWT with golang-jwt/jwt/v5
- **Password Hashing**: bcrypt
- **Documentation**: Swagger with swaggo/fiber-swagger
- **Middleware**: CORS, Logger

## Environment Setup

1. Go 1.19+
2. Required packages:
   - github.com/gofiber/fiber/v2
   - gorm.io/gorm
   - gorm.io/driver/sqlite
   - github.com/golang-jwt/jwt/v5
   - golang.org/x/crypto/bcrypt
   - github.com/swaggo/fiber-swagger