package main

import (
    "fmt"
		"strings"

    refs "github.com/vx6fid/git-go/internal/ref"
    "github.com/vx6fid/git-go/internal/object"
)

func commitCmd(args []string) {
    if len(args) < 2 || args[0] != "-m" {
        fmt.Println("Usage: ugit commit -m <message>")
        return
    }

		message := strings.Join(args[1:], " ")
    author := "Achal <achal@example.com>" // hardcoded, later load from config

    sha, err := object.WriteCommit(message, author, nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

		refPath, isSymbolic, _ := refs.HeadTarget()
		if isSymbolic {
				refs.WriteRef(refPath, sha)
		} else {
				refs.UpdateHEAD(sha)
		}
    fmt.Println(sha)
		fmt.Printf("HEAD target: %s (symbolic: %v)\n", refPath, isSymbolic)
}

