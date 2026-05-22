# Gin User Service

A lightweight RESTful API service built with Go and the Gin web framework for managing user information. This service provides endpoints to create and retrieve user data with a clean architecture pattern.

## 📋 Project Overview

**Gin User Service** is a production-ready microservice that demonstrates best practices in Go development including:
- Clean architecture with clear separation of concerns (Controllers, Services, Models)
- Environment-based configuration management
- Graceful server shutdown
- JSON request/response handling
- Modular code organization

**Module:** `github.com/devdigvijay/gin-user-service`  
**Go Version:** 1.25.5

---

## 🏗️ Architecture & Project Structure

```
gin-user-service/
├── go.mod                                    # Go module definition
├── src/
│   ├── main.go                              # Application entry point
│   ├── controllers/
│   │   └── UserController.go                # HTTP route handlers
│   ├── services/
│   │   └── UserService.go                   # Business logic layer
│   ├── models/
│   │   ├── requests/
│   │   │   └── UserRequest.go               # Request DTOs
│   │   └── responses/
│   │       └── UserResponse.go              # Response DTOs
│   ├── environment/
│   │   ├── environment.configuration.go     # Config loader
│   │   └── configuration/
│   │       ├── environment.dev.yaml         # Development config
│   │       └── environment.prod.yaml        # Production config
│   ├── middlewares/                         # (Placeholder for middleware)
│   └── utils/
│       ├── utils.go                         # Utility functions
│       └── request.json                     # Sample request data
└── README.md                                # Documentation
```

### Architecture Layers

1. **Controllers Layer** (`controllers/`)
   - Handles HTTP routing and request parsing
   - Delegates business logic to services
   - Manages route initialization

2. **Services Layer** (`services/`)
   - Contains core business logic
   - Implements handler functions that return Gin middleware
   - Processes requests and generates responses

3. **Models Layer** (`models/`)
   - Defines Request DTOs (Data Transfer Objects) for incoming data
   - Defines Response DTOs for outgoing data
   - Uses JSON tags for serialization

4. **Environment Layer** (`environment/`)
   - Loads configuration from YAML files
   - Supports environment-specific configs (dev, prod)
   - Integrates environment variables with defaults

5. **Utils Layer** (`utils/`)
   - Helper functions for JSON conversion
   - Environment flag loading

---

## 🔧 Dependencies

### Direct Dependencies
- **`github.com/gin-gonic/gin v1.11.0`** - Web framework for building REST APIs
- **`github.com/ilyakaznacheev/cleanenv v1.5.0`** - Configuration management from YAML and env files

### Key Transitive Dependencies
- **`google.golang.org/protobuf`** - Protocol Buffers support
- **`gopkg.in/yaml.v3`** - YAML parsing
- **`github.com/go-playground/validator/v10`** - Data validation
- **`golang.org/x/crypto`** - Cryptographic utilities
- Plus 30+ other dependencies managed by Go modules

---

## 🚀 Getting Started

### Prerequisites
- Go 1.25.5 or higher
- Basic knowledge of REST APIs and Go

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/devdigvijay/gin-user-service.git
   cd gin-user-service
   ```

2. **Download dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   # Run with default development environment
   go run src/main.go
   
   # Run with specific environment
   go run src/main.go -env=prod
   ```

### Building for Production

```bash
go build -o gin-user-service ./src/main.go
./gin-user-service -env=prod
```

---

## 📡 API Endpoints

All endpoints are prefixed with `/user/`

### 1. Get User Information
- **Endpoint:** `GET /user/`
- **Description:** Retrieves user information
- **Response:**
  ```json
  {
    "Id": "101",
    "UserName": "digv1705",
    "IsActive": true
  }
  ```

### 2. Save User Information
- **Endpoint:** `POST /user/save`
- **Description:** Creates/saves new user information
- **Request Body:**
  ```json
  {
    "firstName": "Digvijay",
    "middleName": "Kumar",
    "lastName": "Singh",
    "userName": "digv1705",
    "Age": 28,
    "Password": "secure_password",
    "isActive": true
  }
  ```
- **Response:**
  ```json
  {
    "Id": "101",
    "UserName": "digv1705",
    "IsActive": true
  }
  ```
- **Status Codes:**
  - `200 OK` - User saved successfully
  - `400 Bad Request` - Invalid request body

---

## ⚙️ Configuration

The application uses YAML-based environment configuration with support for environment variable overrides.

### Configuration Structure

```yaml
app:
  name: gin-user-service
  env: dev                    # Environment: dev or prod
  port: ":8080"              # Server port

server:
  readTimeout: 5s            # HTTP read timeout
  writeTimeout: 10s          # HTTP write timeout
  ShutDownTimeout: 5s        # Graceful shutdown timeout

database:
  host: localhost            # Database host
  port: 5432                 # Database port
  name: users_db             # Database name
  username: admin            # Database username
  password: admin123         # Database password
  ssl: false                 # SSL encryption
```

### Environment Files

- **`environment.dev.yaml`** - Development configuration
  - Port: 8080
  - SSL: Disabled
  - Default credentials for local development

- **`environment.prod.yaml`** - Production configuration (can be customized)

### Environment Variable Overrides

All configuration values can be overridden using environment variables:

```bash
APP_NAME=my-service \
APP_ENV=prod \
APP_PORT=":3000" \
SERVER_READ_TIMEOUT=3s \
SERVER_WRITE_TIMEOUT=8s \
SERVER_SHUTDOWN_TIMEOUT=10 \
DATABASE_HOST=db.example.com \
DATABASE_PORT=5432 \
DATABASE_NAME=prod_db \
DATABASE_USERNAME=admin \
DATABASE_PASSWORD=secure_pwd \
DATABASE_SSL=true \
go run src/main.go -env=prod
```

---

## 🔄 Application Flow

1. **Startup** (`main.go`)
   - Loads environment configuration based on `-env` flag (default: dev)
   - Initializes Gin engine with default middleware
   - Creates UserController instance

2. **Route Initialization** (`UserController.Initialize()`)
   - Registers routes under `/user/` group
   - Maps handler functions from UserService

3. **Request Processing**
   - HTTP request arrives at endpoint
   - UserService handler processes request
   - Response is serialized to JSON and returned

4. **Graceful Shutdown**
   - Listens for SIGINT/SIGTERM signals
   - Gracefully shuts down server within timeout
   - Logs shutdown status

---

## 📊 Data Models

### CreateUserRequest
```go
type CreateUserRequest struct {
    FirstName  string `json:"firstName"`
    MiddleName string `json:"middleName"`
    LastName   string `json:"lastName"`
    UserName   string `json:"userName"`
    Age        int    `json:"Age"`
    Password   string `json:"Password"`
    IsActive   bool   `json:"isActive"`
}
```

### CreateUserResponse
```go
type CreateUserResponse struct {
    Id       string `json:"Id"`
    UserName string `json:"UserName"`
    IsActive bool   `json:"IsActive"`
}
```

---

## 🛠️ Key Features

✅ **Clean Architecture** - Well-organized code with clear separation of concerns  
✅ **Environment Management** - YAML-based configuration with environment variable overrides  
✅ **Graceful Shutdown** - Proper signal handling and server cleanup  
✅ **Modular Design** - Easy to extend with new controllers and services  
✅ **JSON Serialization** - Built-in request/response JSON handling  
✅ **Error Handling** - Proper HTTP status codes and error messages  

---

## 🚦 Running the Application

### Development Mode
```bash
go run src/main.go
# Server runs on http://localhost:8080
```

### Production Mode
```bash
go run src/main.go -env=prod
# Uses configuration from environment.prod.yaml
```

### Testing Endpoints

**Get User:**
```bash
curl http://localhost:8080/user/
```

**Create User:**
```bash
curl -X POST http://localhost:8080/user/save \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Digvijay",
    "middleName": "Kumar",
    "lastName": "Singh",
    "userName": "digv1705",
    "Age": 28,
    "Password": "secure_password",
    "isActive": true
  }'
```

---

## 📝 Logging

The application uses Go's standard `log` package for logging important events:
- Server startup
- Configuration loading errors
- Server shutdown events

Example output:
```
---> Server started on :8080
---> Server exited gracefully
```

---

## 🔐 Security Considerations

⚠️ **Current Implementation Notes:**
- Passwords are currently logged in plain text (development only)
- No authentication/authorization implemented yet
- No database integration (mock responses only)
- All responses are hardcoded (for demo purposes)

**Recommended for Production:**
1. Implement proper password hashing (bcrypt)
2. Add JWT authentication
3. Integrate with a real database (PostgreSQL, MongoDB)
4. Remove sensitive data logging
5. Add input validation and sanitization
6. Implement rate limiting
7. Add HTTPS/TLS support

---

## 📈 Future Enhancements

- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] User authentication with JWT
- [ ] Input validation and error handling
- [ ] Unit and integration tests
- [ ] API documentation with Swagger
- [ ] Middleware for logging, authentication, CORS
- [ ] Docker support
- [ ] CI/CD pipeline
- [ ] Database migrations
- [ ] Rate limiting and throttling

---

## 📄 Module Information

- **Module Name:** `github.com/devdigvijay/gin-user-service`
- **Go Version Required:** 1.25.5+
- **License:** (Add your license here)

---

## 🤝 Contributing

Contributions are welcome! Please ensure:
1. Code follows Go best practices
2. All endpoints are tested
3. Configuration is properly documented
4. Backward compatibility is maintained

---

## ❓ FAQ

**Q: How do I change the server port?**  
A: Set the environment variable `APP_PORT` or modify the YAML configuration file.

**Q: Can I run this without Go installed?**  
A: You can build a binary first: `go build -o gin-user-service ./src/main.go`, then run it: `./gin-user-service`

**Q: How do I debug the application?**  
A: Use `delve` debugger: `dlv debug ./src/main.go` or add logging statements.

**Q: Is the database connection required?**  
A: Currently, the application doesn't connect to a database. It returns hardcoded responses. Database integration needs to be implemented.

---

## 📞 Support

For issues, questions, or suggestions, please open an issue in the repository or contact the maintainer.

---

**Created:** 2026  
**Last Updated:** May 22, 2026  
**Author:** Digvijay
