package main

import (
        "fmt"
        "os"
)
var Version string
func main() {
        if version == "" {
                Version = "0.0.0-dev"
        }

        fmt.Println("Version: ", Version)
}
