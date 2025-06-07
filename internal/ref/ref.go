package ref

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var GitDir = ".git"

func HeadPath() string {
	return filepath.Join(GitDir, "HEAD")
}

func RefsPath(name string) string {
	return filepath.Join(GitDir, "refs", name)
}

// Read HEAD -> return commit SHA or symbolic ref
func ReadHEAD() (string, error) {
	data, err := os.ReadFile(HeadPath())
	if err != nil {
		return "", err
	}

	line := strings.TrimSpace(string(data))
	if strings.HasPrefix(line, "ref: "){
		ref := strings.TrimPrefix(line, "ref: ")
		return ReadRef(ref)
	}
	return line, nil // direct SHA (detached HED)
}

// Read a ref (e.g., refs/heads/master)
func ReadRef(name string) (string, error) {
	path := filepath.Join(GitDir, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err 
	}
	return strings.TrimSpace(string(data)), nil
}

// Write a new commit SHA to a ref (or HEAD)
func WriteRef(name, value string) error {
	path := filepath.Join(GitDir, name)
	os.MkdirAll(filepath.Dir(path), 0755)

	err := os.WriteFile(path, []byte(value+"\n"), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("[ggit] Updated ref %s → %s\n", name, value)
	return nil
}


// Update HEAD (can be symbolic or SHA)
func UpdateHEAD(value string) error {
	return os.WriteFile(HeadPath(), []byte(value), 0644)
}

func HeadTarget() (string, bool, error) {
	data, err := os.ReadFile(HeadPath())
	if err != nil {
		return "", false, err
	}
	line := strings.TrimSpace(string(data))
	if strings.HasPrefix(line, "ref: ") {
		return strings.TrimPrefix(line, "ref: "), true, nil
	}
	return line, false, nil // detached
}
