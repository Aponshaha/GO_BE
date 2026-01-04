# Go Language Guide for Beginners

## 📦 What is a Package in Go?

A **package** in Go is like a folder that contains related code. Think of it as a module or namespace.

### Key Concepts:

1. **Package Declaration**: Every Go file starts with `package <name>`

   ```go
   package models  // This file belongs to the "models" package
   ```

2. **Package = Folder**: All files in the same folder must have the same package name

   - `internal/models/models.go` → `package models`
   - `internal/models/user.go` → `package models` (same package!)

3. **Importing Packages**: Use `import` to use code from other packages

   ```go
   import "ecom/internal/models"  // Import the models package
   ```

4. **Public vs Private**:

   - **Public** (exported): Starts with capital letter → `User`, `GetUser()`
   - **Private** (unexported): Starts with lowercase → `user`, `getUser()`

   Only public functions/types can be used outside the package!

---

## 🏗️ 1. MODELS (`internal/models/models.go`)

### What are Models?

**Models** are data structures that represent your business entities (like User, Product). They define the shape of your data.

### Code Breakdown:

```go
package models  // This file is part of the "models" package

import "time"   // Import time package for timestamps

// User represents a user in the system
type User struct {  // "struct" = a collection of fields (like a class in other languages)
	ID        int       `json:"id"`         // Field with JSON tag
	Email     string    `json:"email"`      // string type
	FirstName string    `json:"first_name"` // JSON tag maps to "first_name"
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"` // time.Time = timestamp
	UpdatedAt time.Time `json:"updated_at"`
}
```

### Explanation:

1. **`type User struct`**:

   - Creates a new type called `User`
   - `struct` = a collection of fields (like a class without methods)
   - Similar to a class in Java/Python, but simpler

2. **Fields**:

   - `ID int` = integer field
   - `Email string` = text field
   - `CreatedAt time.Time` = timestamp field

3. **JSON Tags** (backticks):

   ```go
   `json:"id"`  // When converting to JSON, use "id" as the key
   ```

   - When you convert this struct to JSON, it becomes:
     ```json
     {
       "id": 1,
       "email": "user@example.com",
       "first_name": "John",
       "last_name": "Doe"
     }
     ```

4. **Why Models?**
   - Define data structure once
   - Type safety (Go catches errors at compile time)
   - Reusable across the application

### Example Usage:

```go
// Create a new user
user := models.User{
    ID:        1,
    Email:     "john@example.com",
    FirstName: "John",
    LastName:  "Doe",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}

// Access fields
fmt.Println(user.Email)  // Output: john@example.com
```

---

## 🗄️ 2. REPOSITORIES (`internal/repositories/repositories.go`)

### What is a Repository?

A **Repository** is a design pattern that handles all database operations. It's like a data access layer - it talks to the database and returns your models.

### Why Use Repositories?

- **Separation of Concerns**: Database code is separate from business logic
- **Testability**: Easy to mock for testing
- **Reusability**: Same repository can be used by multiple services

### Code Breakdown:

```go
package repositories

import (
	"database/sql"              // Go's database package
	"ecom/internal/database"  // Our database connection
	"ecom/internal/models"     // Our models
	"fmt"                      // Formatting/errors
)

// UserRepository handles user-related database operations
type UserRepository struct {  // Struct to hold database connection
	db *sql.DB  // Pointer to database connection
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),  // Get the database connection
	}
}
```

### Explanation:

1. **`type UserRepository struct`**:

   - A struct that holds the database connection
   - Methods on this struct can access the database

2. **`NewUserRepository()`**:

   - This is a **constructor function** (convention: starts with `New`)
   - Creates and returns a new `UserRepository` instance
   - Sets up the database connection

3. **Method on Struct**:
   ```go
   func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
       // r = receiver (like "this" or "self" in other languages)
       // *UserRepository = pointer to UserRepository
       // Returns: pointer to User, and error
   }
   ```

### The GetUserByID Method:

```go
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	// SQL query with placeholder $1 (PostgreSQL style)
	query := `SELECT id, email, first_name, last_name, created_at, updated_at
	          FROM users WHERE id = $1`

	// Create empty User struct
	user := &models.User{}  // & = pointer, creates new instance

	// Execute query and scan results into user struct
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,        // & = address of field (where to put the data)
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	// Handle errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")  // User doesn't exist
		}
		return nil, err  // Some other error
	}

	return user, nil  // Success: return user and no error
}
```

### Key Go Concepts:

1. **Pointers (`*` and `&`)**:

   - `*models.User` = pointer to User (reference)
   - `&user.ID` = address of user.ID (where to write data)
   - Pointers avoid copying large structs

2. **Error Handling**:

   ```go
   func() (result, error)  // Go functions often return (result, error)
   ```

   - Always check errors!
   - `nil` means "no error"

3. **QueryRow().Scan()**:
   - `QueryRow()` = execute SQL query (expects 1 row)
   - `.Scan()` = copy database columns into struct fields
   - Order matters! Must match SELECT order

### Example Usage:

```go
// Create repository
repo := repositories.NewUserRepository()

// Get user by ID
user, err := repo.GetUserByID(1)
if err != nil {
    log.Fatal(err)  // Handle error
}

fmt.Println(user.Email)  // Use the user
```

---

## 🔄 3. MIDDLEWARE (`internal/middleware/middleware.go`)

### What is Middleware?

**Middleware** is code that runs BEFORE and AFTER your main handler function. It's like a filter or interceptor.

### Common Uses:

- Logging requests
- Authentication
- CORS headers
- Rate limiting
- Request validation

### Code Breakdown:

```go
package middleware

import (
	"log"       // Logging
	"net/http"  // HTTP server
	"time"      // Time calculations
)

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()  // Record start time

		// Wrap ResponseWriter to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,  // Default 200
		}

		// Call the next handler (your actual endpoint)
		next.ServeHTTP(wrapped, r)

		// After handler runs, log the request
		duration := time.Since(start)
		log.Printf(
			"%s %s %d %v",  // Format: method, path, status, duration
			r.Method,       // GET, POST, etc.
			r.URL.Path,     // /api/users
			wrapped.statusCode,  // 200, 404, etc.
			duration,       // How long it took
		)
	})
}
```

### How Middleware Works:

```
Request comes in
    ↓
Middleware 1 (Logging) - BEFORE
    ↓
Middleware 2 (CORS) - BEFORE
    ↓
Your Handler (actual endpoint)
    ↓
Middleware 2 (CORS) - AFTER
    ↓
Middleware 1 (Logging) - AFTER
    ↓
Response sent
```

### Step-by-Step:

1. **Function Signature**:

   ```go
   func LoggingMiddleware(next http.Handler) http.Handler
   ```

   - Takes: `next` handler (the actual endpoint)
   - Returns: A new handler (wrapped version)

2. **`http.Handler` Interface**:

   ```go
   type Handler interface {
       ServeHTTP(ResponseWriter, *Request)
   }
   ```

   - Any function that implements `ServeHTTP` is a handler

3. **`http.HandlerFunc`**:

   - Converts a function to a Handler
   - `HandlerFunc(func(w, r) { ... })` → Handler

4. **The Flow**:
   ```go
   start := time.Now()              // 1. Record start time
   next.ServeHTTP(wrapped, r)      // 2. Call next handler (your endpoint)
   duration := time.Since(start)    // 3. Calculate duration
   log.Printf(...)                 // 4. Log the request
   ```

### CORS Middleware:

```go
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers (allows browser requests from other domains)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests (browser checks if request is allowed)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return  // Don't call next handler
		}

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}
```

### Custom ResponseWriter:

```go
type responseWriter struct {
	http.ResponseWriter  // Embed original ResponseWriter
	statusCode int       // Store status code
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code              // Save status code
	rw.ResponseWriter.WriteHeader(code) // Call original method
}
```

**Why?** The original `ResponseWriter` doesn't let us read the status code. We wrap it to capture it for logging.

---

## 🔗 How Everything Connects

### Request Flow:

```
1. HTTP Request arrives
   ↓
2. Middleware (Logging, CORS)
   ↓
3. Routes → Handler
   ↓
4. Handler → Service
   ↓
5. Service → Repository
   ↓
6. Repository → Database
   ↓
7. Database returns data
   ↓
8. Repository → Model
   ↓
9. Service → Model
   ↓
10. Handler → JSON Response
   ↓
11. Middleware (Logging)
   ↓
12. HTTP Response sent
```

### Example: Getting a User

```go
// 1. Handler receives request
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := 1  // Extract from request

    // 2. Call service
    user, err := h.userService.GetUser(id)

    // 3. Service calls repository
    // (inside UserService.GetUser)
    return s.userRepo.GetUserByID(id)

    // 4. Repository queries database
    // (inside UserRepository.GetUserByID)
    r.db.QueryRow(query, id).Scan(&user...)

    // 5. Data flows back up
    // Repository → Service → Handler

    // 6. Handler sends JSON response
    json.NewEncoder(w).Encode(user)
}
```

---

## 📚 Go Language Basics

### 1. Variables:

```go
var name string = "John"     // Explicit type
name := "John"               // Type inference (shorthand)
```

### 2. Functions:

```go
func add(a int, b int) int {  // Returns int
    return a + b
}

func divide(a, b int) (int, error) {  // Returns two values
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### 3. Structs:

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "John", Age: 30}
```

### 4. Methods:

```go
func (p *Person) GetName() string {  // Method on Person
    return p.Name
}
```

### 5. Pointers:

```go
x := 10
ptr := &x      // Get address of x
*ptr = 20      // Change value at address
fmt.Println(x) // 20
```

### 6. Error Handling:

```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use result
```

### 7. Interfaces:

```go
type Writer interface {
    Write([]byte) (int, error)
}
// Any type with Write() method implements Writer
```

---

## 🎯 Summary

| Component        | Purpose                     | Example                                   |
| ---------------- | --------------------------- | ----------------------------------------- |
| **Models**       | Data structures             | `User`, `Product`                         |
| **Repositories** | Database operations         | `GetUserByID()`, `GetAllProducts()`       |
| **Services**     | Business logic              | `GetUser()`, `GetAllProducts()`           |
| **Handlers**     | HTTP endpoints              | `GetUser()`, `HealthCheck()`              |
| **Middleware**   | Request/response processing | `LoggingMiddleware()`, `CORSMiddleware()` |

### Key Takeaways:

1. **Packages** = Folders with related code
2. **Models** = Data structures (like classes)
3. **Repositories** = Database access layer
4. **Services** = Business logic layer
5. **Handlers** = HTTP request handlers
6. **Middleware** = Code that runs before/after handlers

This architecture keeps your code:

- ✅ Organized
- ✅ Testable
- ✅ Maintainable
- ✅ Scalable
