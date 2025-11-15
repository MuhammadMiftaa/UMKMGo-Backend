# SapaUMKM Backend - Testing Guide

## 📋 Table of Contents

- [Quick Start](#quick-start)
- [Environment Setup](#environment-setup)
- [Running the Application](#running-the-application)
- [API Testing](#api-testing)
- [Unit Testing](#unit-testing)
- [Database Migration](#database-migration)

---

## 🚀 Quick Start

### Prerequisites

- Go 1.24.4 or higher
- PostgreSQL 13+
- Redis 6+
- Goose (migration tool)

### Installation

1. **Clone the repository**

```bash
git clone <repository-url>
cd UMKMGo-backend
```

2. **Install dependencies**

```bash
go mod download
```

3. **Setup environment variables**

```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Install Goose (if not installed)**

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

---

## 🔧 Environment Setup

### Configure `.env` file

```bash
# Server Configuration
MODE=development
PORT=8080
JWT_SECRET_KEY=your_super_secret_jwt_key

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=sapaumkm_db

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379

# SMTP Configuration
ZOHO_SMTP_HOST=smtp.zoho.com
ZOHO_SMTP_PORT=587
ZOHO_SMTP_USER=your_email@zohomail.com
ZOHO_SMTP_PASSWORD=your_app_password
ZOHO_SMTP_SECURE=true
ZOHO_SMTP_AUTH=true
```

### Database Setup

1. **Create database**

```bash
createdb sapaumkm_db
```

2. **Run migrations**

```bash
make migrate
```

3. **Check migration status**

```bash
make migration-status
```

4. **Rollback migration (if needed)**

```bash
make rollback
```

---

## 🏃 Running the Application

### Start the server

```bash
# Using Make
make api

# Or directly with Go
go run ./cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Available Make Commands

```bash
make api                    # Run the API server
make migrate                # Run all pending migrations
make migrate to=VERSION     # Migrate to specific version
make migration name=NAME    # Create new migration
make rollback               # Rollback last migration
make rollback to=VERSION    # Rollback to specific version
make migration-status       # Check migration status
make seed                   # Run database seeders
make seed-reset            # Reset and rerun seeders
```

---

## 🧪 API Testing

### Using request.http (VS Code REST Client)

1. **Install REST Client extension** in VS Code
2. **Open `request.http`** file
3. **Update the `@token` variable** after login
4. **Click "Send Request"** above each request

### Using Postman

1. **Import collection**

   - Open Postman
   - Click Import
   - Select `postman_collection.json`

2. **Setup environment**

   - Create new environment
   - Add variable: `url` = `http://localhost:8080`
   - Add variable: `token` = (will be set automatically after login)

3. **Test the API**
   - Start with "Authentication > Web Dashboard > Login"
   - Token will be saved automatically
   - Test other endpoints

### Quick Test Endpoints

```bash
# Health check
curl http://localhost:8080/test

# Login (get token)
curl -X POST http://localhost:8080/v1/webauth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"superadmin@example.com","password":"admin123"}'

# Get users (with auth)
curl http://localhost:8080/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN"

# Get programs
curl http://localhost:8080/v1/programs/ \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 🔬 Unit Testing

### Run All Tests

```bash
go test ./internal/service/... -v
```

### Run Specific Test File

```bash
# Test Users Service
go test ./internal/service -run TestRegister -v
go test ./internal/service -run TestLogin -v

# Test Programs Service
go test ./internal/service -run TestCreateProgram -v
go test ./internal/service -run TestGetAllPrograms -v
```

### Run Tests with Coverage

```bash
# Generate coverage report
go test ./internal/service/... -cover

# Generate detailed coverage report
go test ./internal/service/... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### Test Structure

```
internal/service/
├── users.go              # Users service implementation
├── users_test.go         # Users service tests
├── programs.go           # Programs service implementation
└── programs_test.go      # Programs service tests
```

### Unit Test Coverage

**Users Service Tests:**

- ✅ Register (7 test cases)
  - Valid registration
  - Missing required fields
  - Invalid email format
  - Password validations (length, letter, digit)
  - Password mismatch
- ✅ Login (4 test cases)
  - Valid login
  - Missing credentials
  - User not found
  - Wrong password
- ✅ GetUserByID (2 test cases)
- ✅ UpdateUser (2 test cases)
- ✅ DeleteUser (2 test cases)
- ✅ GetListPermissions (1 test case)
- ✅ GetListRolePermissions (1 test case)
- ✅ UpdateRolePermissions (2 test cases)

**Programs Service Tests:**

- ✅ CreateProgram (5 test cases)
  - Valid training program
  - Missing required fields
  - Invalid type
  - Invalid training type
  - Invalid creator user
- ✅ GetAllPrograms (1 test case)
- ✅ GetProgramByID (2 test cases)
- ✅ UpdateProgram (3 test cases)
- ✅ DeleteProgram (2 test cases)
- ✅ ActivateProgram (2 test cases)
- ✅ DeactivateProgram (2 test cases)

### Writing New Tests

Example test structure:

```go
func TestYourFunction(t *testing.T) {
    // Setup
    mockRepo := newMockRepository()
    service := NewService(mockRepo)

    tests := []struct {
        name        string
        input       YourInput
        expectError bool
        errorMsg    string
    }{
        {
            name: "Valid input",
            input: YourInput{...},
            expectError: false,
        },
        {
            name: "Invalid input",
            input: YourInput{...},
            expectError: true,
            errorMsg: "expected error message",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := service.YourFunction(tt.input)

            if tt.expectError {
                if err == nil {
                    t.Errorf("Expected error but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
            }
        })
    }
}
```

---

## 📊 API Endpoints

### Authentication

- `POST /v1/webauth/login` - Web dashboard login
- `POST /v1/webauth/register` - Web dashboard register
- `POST /v1/mobileauth/login` - Mobile app login
- `POST /v1/mobileauth/register` - Mobile app register
- `POST /v1/mobileauth/send/otp` - Send OTP
- `POST /v1/mobileauth/verify/otp` - Verify OTP

### Users Management (Requires Auth)

- `GET /v1/users` - Get all users
- `GET /v1/users/:id` - Get user by ID
- `PUT /v1/users/:id` - Update user
- `DELETE /v1/users/:id` - Delete user

### Permissions & Roles (Requires Auth)

- `GET /v1/permissions` - Get all permissions
- `GET /v1/role-permissions` - Get role permissions
- `POST /v1/role-permissions` - Update role permissions

### Programs (Requires Auth)

- `GET /v1/programs/` - Get all programs
- `GET /v1/programs/:id` - Get program by ID
- `POST /v1/programs/` - Create program
- `PUT /v1/programs/:id` - Update program
- `PUT /v1/programs/activate/:id` - Activate program
- `PUT /v1/programs/deactivate/:id` - Deactivate program
- `DELETE /v1/programs/:id` - Delete program

---

## 🐛 Troubleshooting

### Database Connection Issues

```bash
# Check PostgreSQL is running
pg_isready

# Check database exists
psql -l | grep sapaumkm_db

# Reset database
dropdb sapaumkm_db
createdb sapaumkm_db
make migrate
```

### Redis Connection Issues

```bash
# Check Redis is running
redis-cli ping

# Should return: PONG
```

### Migration Issues

```bash
# Check current migration version
make migration-status

# Force version (if needed)
goose -dir ./config/db/migration postgres "your-connection-string" version

# Reset and remigrate
make rollback to=0
make migrate
```

### Port Already in Use

```bash
# Find process using port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>
```

---

## 📝 Development Workflow

1. **Create new feature branch**

```bash
git checkout -b feature/your-feature-name
```

2. **Write code and tests**

```bash
# Implement feature
# Write unit tests
go test ./internal/service/... -v
```

3. **Test API endpoints**

```bash
# Use request.http or Postman
# Test all CRUD operations
```

4. **Run all tests before commit**

```bash
go test ./... -v
```

5. **Commit and push**

```bash
git add .
git commit -m "Add: your feature description"
git push origin feature/your-feature-name
```

---

## 📚 Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [Fiber Framework](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [Goose Migrations](https://github.com/pressly/goose)
- [REST Client VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

---

## 👥 Team

- Backend Developer: Your Team
- Project: SapaUMKM Backend API

## 📄 License

[Your License Here]
