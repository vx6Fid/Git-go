package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func gitInit() error {
	paths := []string{
		".git",
		filepath.Join(".git", "objects"),
		filepath.Join(".git", "refs", "heads"),
	}

	for _, path := range paths {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err);
		}
	}

	headPath := filepath.Join(".git", "HEAD")
	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return fmt.Errorf("failed to write HEAD: %w", err)
	}

	fmt.Println("Initialized empty Git repository in .git/")
	return nil
}
