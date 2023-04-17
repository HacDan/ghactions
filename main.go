package main

import (
        "fmt"
        "os"
)

func main() {
        version := "0.0.0-dev"
        if os.Getenv("VERSION") != "" {
                version = os.Getenv("VERSION")
        }

        fmt.Println("Version: ", version)
}
