package object

import (
	"bytes"
	"encoding/hex"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
)

func ReadTree(treeHash, basePath string) error {
	objType, content, err := ReadObject(treeHash)
	if err != nil {
		return err
	}

	if objType != "tree" {
		return fmt.Errorf("Not a Tree Object: %s", treeHash)
	}

	if err := os.MkdirAll(basePath, 0755); err != nil {
    return err
	}

	i := 0
	for i < len(content) {
		spaceIdx := bytes.IndexByte(content[i:], ' ')
    mode := string(content[i : i+spaceIdx])
    i += spaceIdx + 1

    nullIdx := bytes.IndexByte(content[i:], 0)
    name := string(content[i : i+nullIdx])
    i += nullIdx + 1

    sha := hex.EncodeToString(content[i : i+20])
    i += 20

		fullPath := filepath.Join(basePath, name)
		if mode == "40000" {
			// subdirectory, we have to recurse
			os.MkdirAll(fullPath, 0755)
			if err := ReadTree(sha, fullPath); err != nil {
				return err
			}
		} else {
			// It's a file : blob
			blobType, blobContent, err := ReadObject(sha)
			if err != nil || blobType != "blob" {
				return fmt.Errorf("Invalid Blob object: %s", sha)
			}
			err = ioutil.WriteFile(fullPath, blobContent, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

