package main // Declares that this program belongs to the main package.

import ( // Starts the import section.
    "fmt"      // Imports fmt for printing messages and formatted output.
    "net/http" // Imports net/http to create and run a web server.
) // Ends the import section.

func main() { // Defines the main function, where the program starts running.

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
        // Registers a function to handle HTTP requests sent to the "/" URL.
        // w is used to send a response back to the browser.
        // r contains information about the incoming request.

        fmt.Fprintf(w, "Hello, World!") 
        // Writes "Hello, World!" as the response sent back to the browser.
    }) // Ends the handler function and registration.

    fmt.Println("Server starting on port 8080...")
    // Prints a message in the terminal telling us that the server is starting.

    http.ListenAndServe(":8080", nil)
    // Starts the web server on port 8080.
    // ":8080" means the server listens on port 8080.
    // nil tells Go to use the default HTTP request router.
    
} // Ends the main function.