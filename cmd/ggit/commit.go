// cmd/ugit/commit.go
package main

import (
    "fmt"

    "github.com/vx6fid/git-go/internal/object"
)

func commitCmd(args []string) {
    if len(args) < 2 || args[0] != "-m" {
        fmt.Println("Usage: ugit commit -m <message>")
        return
    }

    message := args[1]
    author := "Achal <achal@example.com>" // hardcoded, later load from config

    sha, err := object.WriteCommit(message, author, nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(sha)
}

