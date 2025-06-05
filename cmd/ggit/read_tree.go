package main

import (
	"fmt"

	"github.com/vx6fid/git-go/internal/object"
)

func readTreeCmd(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: ggit read-tree <tree-sha>")
		return
	}
	err := object.ReadTree(args[0], ".")
  if err != nil {
		fmt.Println("Error:", err)
  }
}
