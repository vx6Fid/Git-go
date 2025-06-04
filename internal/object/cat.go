package object

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func PrintObject(hash string) error {
	objType, content, err := ReadObject(hash)
	if err != nil {
		return err
	}
	
	switch objType {
	case "blob":
		fmt.Print(string(content))
	case "tree":
		return printTree(content)
	case "commit":
        return printCommit(content)
	default:
		return fmt.Errorf("unsupported object type: %s", objType)
	}
	return nil
}

func ReadObject(hash string) (string, []byte, error) {
	if len(hash) != 40 {
		return "", nil, fmt.Errorf("invalid hash: %s", hash)
	}

	dir := filepath.Join(".git", "objects", hash[:2])
	file := filepath.Join(dir, hash[2:])

	data, err := os.ReadFile(file)
	if err != nil {
		return "", nil, fmt.Errorf("read object file: %w", err)
	}

	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", nil, fmt.Errorf("zlib decode error: %w", err)
	}
	defer r.Close()

	raw, err := io.ReadAll(r)
	if err != nil {
		return "", nil, fmt.Errorf("zlib read error: %w", err)
	}

	//Parse header: "blob <len>\0<content>"
	nullIdx := bytes.IndexByte(raw, 0)
	if nullIdx == -1 {
		return "", nil, fmt.Errorf("Invalid object header")
	}

	header := string(raw[:nullIdx])
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("malformed object header")
	}

	objType := parts[0]
	content := raw[nullIdx+1:]

	return objType, content, nil
}

func printTree(content []byte) error {
	i := 0
	for i < len(content) {
		// mode : until first space
		spaceIdx := bytes.IndexByte(content[i:], ' ')
		if spaceIdx == -1 {
			return fmt.Errorf("invalid tree entry")
		}
		mode := string(content[i:i+spaceIdx])
		i += spaceIdx + 1

		// name until null byte
		nullIdx := bytes.IndexByte(content[i:], 0)
		if nullIdx == -1 {
			return fmt.Errorf("invalid tree entry")
		}
		name := string(content[i: i+nullIdx])
		i += nullIdx + 1
		
		// sha next 20 bytes
		if i+20 > len(content) {
			return fmt.Errorf("incomplete sha")
		}
		sha := hex.EncodeToString(content[i: i+20])
		i += 20
		

		shortType := "blob"
		if mode == "40000"{
			shortType = "tree"
		}

		fmt.Printf("%s %s %s\t%s\n", mode, shortType, sha[:6], name)
	}
	return nil
}

func printCommit(content []byte) error {
    lines := bytes.Split(content, []byte{'\n'})
    inMessage := false

    for _, line := range lines {
        if len(line) == 0 {
            inMessage = true
            fmt.Println()
            continue
        }

        if inMessage {
            fmt.Println(string(line)) // commit message
        } else {
            parts := bytes.SplitN(line, []byte(" "), 2)
            if len(parts) != 2 {
                return fmt.Errorf("invalid commit header line: %s", line)
            }
            fmt.Printf("%s %s\n", parts[0], parts[1])
        }
    }

    return nil
}

