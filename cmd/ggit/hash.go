package main

import (
	"fmt"

	"github.com/vx6fid/git-go/internal/object"
)

func hashObjectCmd(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: ugit hash-object <file>")
		return
	}

	hash, err := object.HashAndStoreBlob(args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(hash)
}
