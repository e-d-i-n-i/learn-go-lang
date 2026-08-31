# Hello Go

A simple Go web server that responds with **"Hello, World!"** on port `8080`.

This project is designed as a beginner-friendly introduction to running a Go backend application.

## Prerequisites

Before running the project, make sure you have:

* [Go](https://go.dev/) **1.21 or later**
* [Git](https://git-scm.com/) installed
* A code editor such as VS Code

You can check your Go installation with:

```bash
go version
```

You should see something similar to:

```text
go version go1.21.x windows/amd64
```

## Project Structure

```text
hello-go/
└── main.go
```

* `main.go` — Contains the Go application and HTTP server.

## Running the Application

### 1. Clone the repository

```bash
git clone <repository-url>
```

Move into the project directory:

```bash
cd hello-go
```

### 2. Run the server

You can run the application directly without creating a binary:

```bash
go run main.go
```

You should see the server start on port `8080`.

### 3. Open the application

Open your browser and visit:

```text
http://localhost:8080
```

You should see:

```text
Hello, World!
```

### 4. Stop the server

To stop the running server, press:

```text
Ctrl + C
```

## Building the Application

Instead of running the source code directly, you can compile the application into an executable.

Run:

```bash
go build -o hello-go
```

This creates a binary called:

```text
hello-go
```

On Windows, depending on your environment, you may see:

```text
hello-go.exe
```

## Running the Built Application

### Git Bash

If you are using Git Bash:

```bash
./hello-go
```

On Windows, you can also run:

```bash
./hello-go.exe
```

Then open:

```text
http://localhost:8080
```

## Development Workflow

For beginners, the typical workflow is:

```text
Write code
   ↓
Run with go run
   ↓
Test in browser
   ↓
Make changes
   ↓
Run again
```

For example:

```bash
go run main.go
```

Once the application is ready, build it:

```bash
go build -o hello-go
```

Then run the compiled application:

```bash
./hello-go
```

## Useful Go Commands

| Command                | Purpose                             |
| ---------------------- | ----------------------------------- |
| `go version`           | Check the installed Go version      |
| `go run main.go`       | Run the application                 |
| `go build`             | Compile the application             |
| `go build -o hello-go` | Compile with a specific binary name |
| `go fmt ./...`         | Format Go source code               |
| `go test ./...`        | Run tests                           |

## Troubleshooting

### Port 8080 is already in use

If you see an error indicating that port `8080` is already being used, another application may already be running on that port.

Stop the other application or change the port in `main.go`.

### `go` command not found

If you see:

```text
go: command not found
```

make sure Go is installed and available in your system's `PATH`.

Check with:

```bash
go version
```

If the command still does not work, restart VS Code after installing Go.

## Next Steps

Once you understand this simple server, the next concepts to learn are:

1. Go variables and data types
2. Functions
3. Structs
4. Packages
5. HTTP handlers
6. JSON responses
7. Routing
8. Request validation
9. Middleware
10. Connecting Go to a database

This simple **Hello Go** application is the starting point for building a complete Go backend API.
