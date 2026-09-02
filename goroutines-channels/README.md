# Go Goroutines & Channels: Hello World

A minimal Go project demonstrating the use of goroutines and channels.

## Description

This program launches a goroutine that sends a message through a channel after a short delay. The main goroutine waits to receive the message and then prints it. It's the simplest possible illustration of concurrency and communication in Go.

## Code Overview

- A string channel is created using `make(chan string)`.
- An anonymous goroutine is started with `go func() { ... }()`.
- Inside the goroutine, `time.Sleep(1 * time.Second)` simulates work, then sends `"Hello from goroutine!"` into the channel.
- The main function blocks at `msg := <-messageChannel` until the message is received.
- The message is printed, followed by a completion message.

## Requirements

- Go 1.x installed on your system.

## How to Run

1. Save the code as `main.go`.
2. Open a terminal in the project directory.
3. Run the program:

   ```bash
   go run main.go

## Expected Output

After a one-second delay, you should see:

```
Hello from goroutine!
Main function finished.
```

## Next Steps

Experiment with:
- Adding multiple goroutines sending messages to the same channel.
- Using a buffered channel (`make(chan string, 2)`).
- Using `select` to handle multiple channels.
