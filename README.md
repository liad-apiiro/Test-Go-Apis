# Go API Testing Project

This project contains mock APIs implemented using various Go web frameworks. The purpose is to challenge code analysis tools by including:

- Standard HTTP methods (GET, POST, PUT, DELETE)
- Multiple methods on the same endpoint
- Non-standard/custom HTTP methods (e.g., "FUNKYTOWN", "DANCE", "PARTY")
- Discouraged coding practices (missing method validation, error handling, input validation, etc.)

## Frameworks Used

1. **net/http** (Standard Library) - Port 8080
2. **Gin** - Port 8081
3. **Echo** - Port 8082
4. **Fiber** - Port 8083
5. **Chi** - Port 8084

## Running the Project

```bash
go run .
```

This will start all 5 servers concurrently. Press Ctrl+C to stop all servers.

## API Endpoints

### Standard net/http (Port 8080)

#### Standard Methods
- `GET /api/v1/users` - Get users
- `POST /api/v1/users/create` - Create user
- `GET|POST|PUT|DELETE /api/v1/products` - Multiple methods on same endpoint

#### Non-Standard Methods
- `FUNKYTOWN /api/v1/funkytown` - Custom method
- `DANCE /api/v1/dance` - Custom method
- `FUNKYTOWN|DANCE|PARTY /api/v1/custom` - Multiple custom methods

#### Discouraged Practices
- `ANY /api/v1/bad/no-method-check` - No method validation
- `POST /api/v1/bad/no-error-handling` - Missing error handling
- `GET /api/v1/bad/no-content-type` - Missing Content-Type header
- `POST /api/v1/bad/no-validation` - No input validation

### Gin Framework (Port 8081)

#### Standard Methods
- `GET /api/v2/users` - Get users
- `POST /api/v2/users` - Create user
- `GET|POST|PUT|DELETE /api/v2/products` - Multiple methods

#### Non-Standard Methods
- `FUNKYTOWN /api/v2/funkytown` - Custom method
- `DANCE /api/v2/dance` - Custom method
- `FUNKYTOWN|PARTY /api/v2/custom` - Multiple custom methods

#### Discouraged Practices
- `ANY /api/v2/bad/no-method-check` - Accepts any method
- `POST /api/v2/bad/no-error-handling` - Missing error check on BindJSON
- `POST /api/v2/bad/no-validation` - No input validation
- `ANY /api/v2/bad/*` - Method not allowed but returns 200

### Echo Framework (Port 8082)

#### Standard Methods
- `GET /api/v3/users` - Get users
- `POST /api/v3/users` - Create user
- `GET|POST|PUT|DELETE /api/v3/products` - Multiple methods

#### Non-Standard Methods
- `FUNKYTOWN /api/v3/funkytown` - Custom method
- `DANCE /api/v3/dance` - Custom method
- `FUNKYTOWN|PARTY /api/v3/custom` - Multiple custom methods

#### Discouraged Practices
- `ANY /api/v3/bad/no-method-check` - Accepts any method
- `POST /api/v3/bad/no-error-handling` - Missing error check on Bind
- `POST /api/v3/bad/no-validation` - No input validation
- `GET /api/v3/bad/no-error-return` - Not properly handling errors

### Fiber Framework (Port 8083)

#### Standard Methods
- `GET /api/v4/users` - Get users
- `POST /api/v4/users` - Create user
- `GET|POST|PUT|DELETE /api/v4/products` - Multiple methods

#### Non-Standard Methods
- `FUNKYTOWN /api/v4/funkytown` - Custom method
- `DANCE /api/v4/dance` - Custom method
- `FUNKYTOWN|PARTY /api/v4/custom` - Multiple custom methods

#### Discouraged Practices
- `ALL /api/v4/bad/no-method-check` - Accepts any method
- `POST /api/v4/bad/no-error-handling` - Missing error check on BodyParser
- `POST /api/v4/bad/no-validation` - No input validation
- `GET /api/v4/bad/no-error-return` - Missing error handling

### Chi Framework (Port 8084)

#### Standard Methods
- `GET /api/v5/users` - Get users
- `POST /api/v5/users` - Create user
- `GET|POST|PUT|DELETE /api/v5/products` - Multiple methods

#### Non-Standard Methods
- `FUNKYTOWN /api/v5/funkytown` - Custom method
- `DANCE /api/v5/dance` - Custom method
- `FUNKYTOWN|PARTY /api/v5/custom` - Multiple custom methods

#### Discouraged Practices
- `ANY /api/v5/bad/no-method-check` - Accepts any method (HandleFunc)
- `POST /api/v5/bad/no-error-handling` - Missing error handling for body read
- `POST /api/v5/bad/no-validation` - No input validation
- `GET /api/v5/bad/no-content-type` - Missing Content-Type header

## Testing Non-Standard Methods

To test non-standard HTTP methods, you can use curl:

```bash
# Test FUNKYTOWN method
curl -X FUNKYTOWN http://localhost:8080/api/v1/funkytown

# Test DANCE method
curl -X DANCE http://localhost:8081/api/v2/dance

# Test PARTY method
curl -X PARTY http://localhost:8082/api/v3/custom
```

## Discouraged Practices Included

This project intentionally includes several discouraged coding practices to challenge code analysis tools:

1. **Missing Method Validation**: Endpoints that accept any HTTP method without checking
2. **Missing Error Handling**: Code that ignores errors from operations like body parsing
3. **Missing Input Validation**: Endpoints that process user input without validation
4. **Missing Content-Type Headers**: Responses without proper Content-Type headers
5. **Improper Error Returns**: Handlers that don't properly return or handle errors
6. **Insecure Practices**: Code that echoes user input without sanitization

## Project Structure

```
.
├── main.go              # Main entry point, starts all servers
├── nethttp_server.go    # Standard net/http implementation
├── gin_server.go        # Gin framework implementation
├── echo_server.go       # Echo framework implementation
├── fiber_server.go      # Fiber framework implementation
├── chi_server.go        # Chi framework implementation
├── go.mod              # Go module dependencies
└── README.md           # This file
```

## Dependencies

- `github.com/gin-gonic/gin v1.9.1`
- `github.com/labstack/echo/v4 v4.11.4`
- `github.com/gofiber/fiber/v2 v2.52.0`
- `github.com/go-chi/chi/v5 v5.0.11`

## Notes

- All servers run concurrently in separate goroutines
- Each framework demonstrates different ways to handle HTTP methods
- Non-standard methods are implemented using framework-specific features
- Discouraged practices are clearly marked with `/bad/` in the path

