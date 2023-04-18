package main

import (
        "fmt"
)
var Version string
func main() {
        if Version == "" {
                Version = "0.0.0-dev"
        }

        fmt.Println("Version: ", Version)
}
