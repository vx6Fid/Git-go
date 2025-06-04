package object

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"path/filepath"
)

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

// Returns SHA of tree object created from the current directory
func WriteTree(dir string) (string, error) {
	entries := []TreeEntry{}

	items, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("read dir: %w", err)
	}

	for _, item := range items {
		if item.Name() == ".git" {
			continue // Skipping the .git 
		}

		path := filepath.Join(dir, item.Name())
		info, _ := item.Info()

		if item.IsDir() {
			hash, err := WriteTree(path)
			if err != nil {
				return "", err
			}
			entries = append(entries, TreeEntry{
				Mode: "40000", // directory
				Name: item.Name(),
				Hash: hash,
			})
		} else {
			hash, err := HashAndStoreBlob(path)
			if err != nil {
				return "", err
			}

			mode := fileModeToGitMode(info.Mode())
			entries = append(entries, TreeEntry{
				Mode: mode,
				Name: item.Name(),
				Hash: hash,
			})
		}
	}

	// Sort by Name(for deterministic hashes -- same structure always gives the same tree hash)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	var buf bytes.Buffer
	for _, entry := range entries {
		fmt.Fprintf(&buf, "%s %s\x00", entry.Mode, entry.Name)
		shaBytes := hexToBytes(entry.Hash)
		buf.Write(shaBytes)
	}

	data := buf.Bytes()
	header := fmt.Sprintf("tree %d\x00", len(data))
	full := append([]byte(header), data...)

	sha := fmt.Sprintf("%x", sha1.Sum(full))

	dirPath := filepath.Join(".git", "objects", sha[:2])
	filePath := filepath.Join(dirPath, sha[2:])

	os.MkdirAll(dirPath, 0755)

  var z bytes.Buffer
  w := zlib.NewWriter(&z)
  w.Write(full)
  w.Close()

  os.WriteFile(filePath, z.Bytes(), 0644)

  return sha, nil
}

func fileModeToGitMode(m fs.FileMode) string {
    if m&0111 != 0 {
        return "100755" // executable
    }
    return "100644"     // normal file
}

func hexToBytes(hexStr string) []byte {
    b := make([]byte, 20)
    for i := 0; i < 20; i++ {
        fmt.Sscanf(hexStr[i*2:i*2+2], "%02x", &b[i])
    }
    return b
}
