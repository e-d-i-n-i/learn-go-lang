# Core Backend API Concepts

After learning the fundamentals of Go, the next step is to understand how Go is used to build **backend APIs**.

This section introduces the core concepts needed to build a structured, secure, testable, and production-ready API.

---

## 1. Core API Concepts

These are the fundamental building blocks of a backend API.

| # | Concept                | Beginner Introduction                                                    | What They Should Learn                                               |
| - | ---------------------- | ------------------------------------------------------------------------ | -------------------------------------------------------------------- |
| 1 | **Handlers**           | The part of the API that receives HTTP requests and sends responses.     | GET, POST, PUT, PATCH, DELETE, request/response                      |
| 2 | **Models**             | Structures that represent the data used by the application and database. | Go structs, fields, relationships                                    |
| 3 | **DTOs**               | Objects used to define exactly what data an API accepts or returns.      | Request DTOs, response DTOs, separating API data from models         |
| 4 | **Response Structure** | A consistent format for API responses makes APIs easier to consume.      | `message`, `error`, `data`, pagination                               |
| 5 | **Validation**         | Checking that incoming data is correct before processing it.             | Required fields, email validation, password rules, custom validation |

---

## 2. Structuring the Backend

Once the basic API concepts are understood, we can organize the application into separate layers.

| # | Concept               | Beginner Introduction                                                    | What They Should Learn                                 |
| - | --------------------- | ------------------------------------------------------------------------ | ------------------------------------------------------ |
| 6 | **Repository**        | A layer responsible for communicating with the database.                 | CRUD, database queries, separation of concerns         |
| 7 | **Middleware**        | Code that runs before or after a request reaches the handler.            | Logging, authentication, CORS, error handling          |
| 8 | **Authentication**    | Verifying who a user is before allowing access to protected resources.   | Passwords, login, JWT, tokens, protected routes        |
| 9 | **API Rate Limiting** | Controlling how many requests a client can make within a period of time. | Request limits, IP/user-based limits, preventing abuse |

---

## 3. Database, Documentation & Testing

After understanding how the application is structured, we introduce the tools and practices used to make the API reliable and production-ready.

| #  | Concept                   | Beginner Introduction                                                   | What They Should Learn                                   |
| -- | ------------------------- | ----------------------------------------------------------------------- | -------------------------------------------------------- |
| 10 | **Migration**             | A controlled way to create and change the database structure.           | Tables, columns, indexes, up/down migrations             |
| 11 | **Docker**                | Packaging the application and its dependencies so it runs consistently. | Dockerfile, images, containers, Docker Compose           |
| 12 | **Swagger Documentation** | Documenting and testing API endpoints through an interactive interface. | OpenAPI, endpoint documentation, Swagger UI              |
| 13 | **Automated Testing**     | Automatically checking that the API behaves correctly.                  | Unit tests, integration tests, HTTP handler tests, mocks |

---

# 4. How These Concepts Work Together

A typical request can flow through several layers of the backend:

```text
Client
  │
  │ HTTP Request
  ▼
Middleware
  │
  │ Authentication / Logging / Rate Limiting
  ▼
Handler
  │
  │ Validate Request
  ▼
DTO
  │
  │ Business Logic
  ▼
Repository
  │
  │ Database Query
  ▼
Database
  │
  │ Data
  ▼
Repository
  │
  ▼
Handler
  │
  │ Response DTO
  ▼
Client
```

For example, when a user creates a device:

```text
POST /api/v1/devices
        │
        ▼
   Middleware
        │
        ▼
     Handler
        │
        ▼
   Validate DTO
        │
        ▼
   Repository
        │
        ▼
    Database
        │
        ▼
    Response
```

This separation makes the application easier to:

* Understand
* Test
* Maintain
* Debug
* Extend
* Scale

---

# 5. Example Backend Structure

As the project grows, we can organize the Go application into different packages:

```text
project/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── dto/
│   ├── repositories/
│   ├── middleware/
│   ├── services/
│   └── auth/
│
├── migrations/
│
├── docs/
│
├── tests/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

Each part has a specific responsibility:

| Folder          | Responsibility                      |
| --------------- | ----------------------------------- |
| `cmd/`          | Application entry points            |
| `handlers/`     | HTTP request handling               |
| `models/`       | Database/application models         |
| `dto/`          | API request and response structures |
| `repositories/` | Database operations                 |
| `middleware/`   | Request processing middleware       |
| `services/`     | Business logic                      |
| `auth/`         | Authentication and authorization    |
| `migrations/`   | Database schema changes             |
| `docs/`         | API documentation                   |
| `tests/`        | Automated tests                     |

---

# 6. Learning Order

A good learning sequence is to introduce the concepts gradually:

```text
Go Fundamentals
      │
      ▼
HTTP & REST APIs
      │
      ▼
Handlers
      │
      ▼
Models & DTOs
      │
      ▼
Validation
      │
      ▼
Response Structures
      │
      ▼
Repositories & Database
      │
      ▼
Middleware
      │
      ▼
Authentication
      │
      ▼
Rate Limiting
      │
      ▼
Database Migrations
      │
      ▼
Swagger Documentation
      │
      ▼
Automated Testing
      │
      ▼
Docker & Deployment
```

---

# 7. What Beginners Should Be Able to Build

By the end of these topics, beginners should be able to build a basic Go REST API with:

* HTTP endpoints
* Request and response DTOs
* Data validation
* Database integration
* Repository pattern
* Middleware
* User authentication
* Protected routes
* API rate limiting
* Database migrations
* Swagger/OpenAPI documentation
* Automated tests
* Docker support

The goal is not just to learn individual concepts, but to understand **how they work together to form a real backend application**.
