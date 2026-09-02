# Two Goroutines, Two Channels

A simple Go program that demonstrates concurrency using two goroutines and two channels. Each goroutine prints its own status and sends a message to its dedicated channel, while the main function waits to receive from both.

## Description

This program shows how multiple goroutines can run concurrently and communicate with the main goroutine through separate channels. It also illustrates that goroutines can print to the console independently, resulting in interleaved output.

## Code Overview

- Two unbuffered string channels are created: `channel1` and `channel2`.
- **Goroutine 1** prints a start message, sleeps for 1 second (simulating work), sends `"Result from goroutine 1"` to `channel1`, then prints a confirmation.
- **Goroutine 2** does the same but sleeps for 2 seconds and uses `channel2`.
- The main goroutine receives from `channel1` and prints the result, then receives from `channel2` and prints that result.
- Finally, it prints a completion message.

## Requirements

- Go 1.x or later installed on your system.

## How to Run

1. Save the code as `main.go`.
2. In a terminal, navigate to the project directory.
3. Run:

   ```bash
   go run main.go
   ```

## Expected Output

The output may vary in order because the goroutines run concurrently. A typical run might look like:

```
Goroutine 1 started, doing work...
Goroutine 2 started, doing work...
Goroutine 1 sent to channel1
Main received: Result from goroutine 1
Goroutine 2 sent to channel2
Main received: Result from goroutine 2
Main function finished.
```

However, because of timing, you might see the messages from goroutine 2 appear after main has already received from channel1, or other valid interleavings. The main function always receives from `channel1` before `channel2`, so the "Main received" lines will appear in that order, but the goroutine prints can appear at different points.

## Key Concepts Demonstrated

- **Goroutines**: Lightweight threads started with the `go` keyword.
- **Channels**: Typed conduits for communication between goroutines.
- **Concurrency**: Both goroutines run independently, printing and sleeping without blocking the main goroutine (except when it waits to receive).
- **Blocking receives**: The main goroutine blocks on `<-channel1` until the value is sent.

## Extend It

- Use a `select` statement to receive from whichever channel is ready first.
- Add more goroutines and channels.
- Experiment with buffered channels (e.g., `make(chan string, 1)`) and observe the behavior.
```