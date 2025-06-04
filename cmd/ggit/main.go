package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("usage: ugit <commands>")
		return 
	}

	switch os.Args[1] {
	case "init":
		if err := gitInit(); err != nil {
			fmt.Println("Error:", err)
		}
	case "hash-object":
		hashObjectCmd(os.Args[2:])
	case "cat-file":
		catFileCmd(os.Args[2:])
	case "write-tree":
    writeTreeCmd()
	case "commit":
    commitCmd(os.Args[2:])
	default:
		fmt.Println("Unknown command: ", os.Args[1]);
	}
}
