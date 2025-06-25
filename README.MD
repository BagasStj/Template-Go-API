# Template Go API

Template Go API adalah sebuah base project untuk pengembangan RESTful API menggunakan Go (Golang) dengan arsitektur clean code dan best practices. 

## 🚀 Features

- **Clean Architecture**: Implementasi clean code dengan separation of concerns
- **RESTful API**: API endpoints yang mengikuti standar REST
- **Database Integration**: Support PostgreSQL dengan read/write database separation
- **CORS Support**: Cross-Origin Resource Sharing configuration
- **Structured Logging**: Menggunakan Zap logger untuk logging yang terstruktur
- **Error Handling**: Centralized error handling dengan custom error types
- **Pagination**: Built-in pagination support untuk list endpoints
- **Validation**: Input validation menggunakan Gin validator
- **Environment Configuration**: Flexible configuration menggunakan environment variables
- **Health Check**: Health check endpoint untuk monitoring

## 📁 Project Structure

```
template-go-api/
├── config/
│   └── config.go              # Configuration management
├── controllers/
│   └── user_controller.go     # HTTP request handlers
├── database/
│   └── connection.go          # Database connection management
├── domains/
│   └── base.go               # Domain models/entities
├── logger/
│   └── logger.go             # Logging interface and implementation
├── models/
│   ├── requests.go           # Request DTOs
│   ├── responses.go          # Response DTOs
│   └── errors/
│       └── errors.go         # Error handling models
├── repositories/
│   └── repository.go         # Data access layer
├── routes/
│   ├── routes.go            # Route configuration
│   └── user_routes.go       # User-specific routes
├── services/
│   └── user_service.go      # Business logic layer
├── utils/
│   ├── pagination.go        # Pagination utilities
│   └── validator.go         # Validation utilities
├── main.go                  # Application entry point
├── go.mod                   # Go module dependencies
└── README.md               # Project documentation
```

## 🏗️ Architecture

Project ini menggunakan **Clean Architecture** dengan layer-layer berikut:

1. **Controllers** - HTTP request handlers, menangani input/output HTTP
2. **Services** - Business logic layer, berisi logika bisnis aplikasi
3. **Repositories** - Data access layer, berinteraksi dengan database
4. **Domains** - Entity models, representasi data dalam aplikasi
5. **Models** - DTOs untuk request/response dan error handling
6. **Utils** - Utility functions yang dapat digunakan di berbagai layer

## 📋 Prerequisites

- Go 1.23.0 atau lebih baru
- PostgreSQL 12+ (opsional, untuk development bisa menggunakan in-memory)
- Git


## Library yang dipakai

### Core Dependencies

#### Web Framework & HTTP
- **[Gin](https://gin-gonic.com/)** (`github.com/gin-gonic/gin v1.10.0`) - High-performance HTTP web framework
- **[CORS](https://github.com/gin-contrib/cors)** (`github.com/gin-contrib/cors v1.7.2`) - Cross-Origin Resource Sharing middleware untuk Gin

#### Database & ORM
- **[GORM](https://gorm.io/)** (`gorm.io/gorm v1.30.0`) - Object Relational Mapping library untuk Go
- **[PostgreSQL Driver](https://github.com/go-gorm/postgres)** (`gorm.io/driver/postgres v1.5.9`) - PostgreSQL driver untuk GORM

#### Configuration & Environment
- **[GoDotEnv](https://github.com/joho/godotenv)** (`github.com/joho/godotenv v1.5.1`) - Untuk load environment variables dari file .env
- **[Env](https://github.com/codingconcepts/env)** (`github.com/codingconcepts/env v0.0.0-20240618133406-5b0845441187`) - Untuk binding environment variables ke struct

#### Logging
- **[Zap](https://go.uber.org/zap)** (`go.uber.org/zap v1.27.0`) - Structured logging library dengan performa tinggi

#### Utilities
- **[UUID](https://github.com/google/uuid)** (`github.com/google/uuid v1.6.0`) - Untuk generate UUID

### Indirect Dependencies (Auto-imported)

#### JSON & Serialization
- **Sonic** (`github.com/bytedance/sonic`) - High-performance JSON library
- **Go-JSON** (`github.com/goccy/go-json`) - Alternative JSON library

#### Validation
- **Validator** (`github.com/go-playground/validator/v10`) - Struct validation library
- **Universal Translator** (`github.com/go-playground/universal-translator`) - Untuk internationalization

#### Database Connection
- **PGX** (`github.com/jackc/pgx/v5`) - PostgreSQL driver dan toolkit
- **Puddle** (`github.com/jackc/puddle/v2`) - Connection pool untuk PostgreSQL

#### System & Platform
- **IsATTY** (`github.com/mattn/go-isatty`) - Untuk deteksi terminal
- **CPUID** (`github.com/klauspost/cpuid/v2`) - CPU feature detection


## ⚙️ Installation

1. **Clone repository:**
   ```bash
   git clone <repository-url>
   cd template-go-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Setup environment variables:**
   ```bash
   cp .env.example .env
   ```
   Edit file `.env` sesuai dengan konfigurasi Anda.

## 🔧 Configuration

Buat file `.env` di root directory dengan konfigurasi berikut:

```env
# Server Configuration
SERVER_PORT=8080
ENV=development

# Database Configuration
DB_HOST_RO=localhost
DB_HOST_RW=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWD=your_db_password
DB_NAME=your_db_name

# CORS Configuration
CORS_ALLOW_ORIGINS=*

# JWT Secrets (optional)
ACCESS_TOKEN_SECRET=your_access_token_secret
REFRESH_TOKEN_SECRET=your_refresh_token_secret

# API Configuration
API_SECRET=your_api_secret

# Basic Auth (optional)
BASIC_USERNAME=admin
BASIC_PASSWORD=admin123
```

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `SERVER_PORT` | No | `8080` | Port aplikasi akan berjalan |
| `ENV` | No | `development` | Environment (development/staging/production) |
| `DB_HOST_RO` | No | - | Database host untuk read operations |
| `DB_HOST_RW` | No | - | Database host untuk write operations |
| `DB_PORT` | No | `5432` | Database port |
| `DB_USER` | No | - | Database username |
| `DB_PASSWD` | No | - | Database password |
| `DB_NAME` | No | - | Database name |
| `CORS_ALLOW_ORIGINS` | No | `*` | Allowed CORS origins |

## 🚦 Running the Application

### Development Mode

```bash
go run main.go
```

### Build and Run

```bash
# Build binary
go build -o template-go-api

# Run binary
./template-go-api
```

## 📡 API Endpoints

### Health Check

```http
GET /health
```

**Response:**
```json
{
  "status": "OK",
  "message": "Template Go API is running"
}
```

### API Version Info

```http
GET /v1/
```

**Response:**
```json
{
  "message": "Welcome to Template Go API v1",
  "version": "1.0.0"
}
```

### Users API

#### Get All Users

```http
GET /v1/users?page=1&limit=10
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)

**Response:**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "name": "John Doe",
      "email": "john@example.com",
      "username": "johndoe",
      "phone_number": "+1234567890",
      "is_active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_records": 50,
    "has_next": true,
    "has_prev": false
  }
}
```

#### Get User by ID

```http
GET /v1/users/{id}
```

**Response:**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": "uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "username": "johndoe",
    "phone_number": "+1234567890",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Create User

```http
POST /v1/users
```

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "username": "johndoe",
  "password": "password123",
  "confirm_password": "password123",
  "phone_number": "+1234567890"
}
```

**Response:**
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "username": "johndoe",
    "phone_number": "+1234567890",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Update User

```http
PUT /v1/users/{id}
```

**Request Body:**
```json
{
  "name": "John Smith",
  "username": "johnsmith",
  "phone_number": "+9876543210"
}
```

#### Delete User

```http
DELETE /v1/users/{id}
```

**Response:**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

## 🧪 Testing

### Unit Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Integration Tests

```bash
# Run integration tests
go test -tags=integration ./...
```

### API Testing

Anda dapat menggunakan tools seperti Postman, Insomnia, atau curl untuk testing API endpoints.

Contoh menggunakan curl:

```bash
# Health check
curl http://localhost:8080/health

# Get users
curl http://localhost:8080/v1/users

# Create user
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "username": "johndoe",
    "password": "password123",
    "confirm_password": "password123"
  }'
```

## 🚀 Deployment

### Manual Deployment

1. **Build aplikasi:**
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
   ```

2. **Upload ke server dan jalankan:**
   ```bash
   ./main
   ```


### Development Guidelines

- Ikuti Go coding standards dan conventions
- Tulis unit tests untuk fungsi-fungsi baru
- Update dokumentasi jika diperlukan
- Gunakan meaningful commit messages
- Pastikan code linter bersih (`go fmt`, `go vet`)

## 🔧 Development Tools

### Recommended Tools

- **Air**: Hot reload untuk development
- **golangci-lint**: Linting tool
- **Postman/Insomnia**: API testing
- **pgAdmin**: PostgreSQL management
- **VS Code**: Code editor dengan Go extension

### VS Code Extensions

- Go (Google)
- REST Client
- GitLens
- Error Lens

## 📚 References

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM ORM](https://gorm.io/)
- [Zap Logger](https://github.com/uber-go/zap)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Support

Jika ada pertanyaan atau masalah, silakan buat issue di repository ini atau hubungi tim development.

---

**Happy Coding! 🎉** 