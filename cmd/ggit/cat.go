package main

import (
	"fmt"
	"github.com/vx6fid/git-go/internal/object"
)

func catFileCmd(args []string) {
	if len(args) < 2 || args[0] != "-p" {
		fmt.Println("Usage: ggit cat-file -p <hash>")
		return
	}

	if err := object.PrintObject(args[1]); err != nil {
		fmt.Println("Error:", err)
	}
}
