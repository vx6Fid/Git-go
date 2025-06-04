package main

import (
    "fmt"
    "github.com/vx6fid/git-go/internal/object"
)

func writeTreeCmd() {
    sha, err := object.WriteTree(".")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(sha)
}
