package main

import (
        "fmt"
        "os"
)
var version string
func main() {
        if version == "" {
                version  "0.0.0-dev"
        }

        fmt.Println("Version: ", version)
}
