package main

import (
    "fmt"
    "time"
)

func main() {

    channel1 := make(chan string)
    channel2 := make(chan string)

    go func() {
        fmt.Println("Goroutine 1 started, doing work...")
        time.Sleep(3 * time.Second)

        channel1 <- "Result from goroutine 1"
        fmt.Println("Goroutine 1 sent to channel1")
    }()

    go func() {
        fmt.Println("Goroutine 2 started, doing work...")
        time.Sleep(1 * time.Second)

        channel2 <- "Result from goroutine 2"
        fmt.Println("Goroutine 2 sent to channel2")
    }()

    // Receive whichever result is ready first.
    for i := 0; i < 2; i++ {
        select {
        case msg := <-channel1:
            fmt.Println("Main received:", msg)

        case msg := <-channel2:
            fmt.Println("Main received:", msg)
        }
    }

    fmt.Println("Main function finished.")
}