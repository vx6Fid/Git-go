// cmd/ugit/init_test.go
package main

import (
    "os"
    "testing"
)

func TestGitInit(t *testing.T) {
    // Use a temporary directory to avoid polluting current working dir
    tmp := t.TempDir()
    os.Chdir(tmp)

    err := gitInit()
    if err != nil {
        t.Fatalf("gitInit failed: %v", err)
    }

    expectedPaths := []string{
        ".git",
        ".git/objects",
        ".git/refs/heads",
        ".git/HEAD",
    }

    for _, path := range expectedPaths {
        if _, err := os.Stat(path); os.IsNotExist(err) {
            t.Errorf("expected %s to exist", path)
        }
    }
}

