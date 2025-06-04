package object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func HashAndStoreBlob(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "",fmt.Errorf("read file: %w", err)
	}

	header := fmt.Sprintf("blob %d\x00", len(content))
	store := append([]byte(header), content...)

	hash := fmt.Sprintf("%x", sha1.Sum(store))

	dir := filepath.Join(".git", "objects", hash[:2])
	file := filepath.Join(dir, hash[2:])

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("mkdir: %w", err);
	}

	var buf bytes.Buffer
	z := zlib.NewWriter(&buf)
	if _, err := z.Write(store); err != nil {
		return "", fmt.Errorf("compress: %w", err)
	}
	z.Close();

	if err := os.WriteFile(file, buf.Bytes(), 0644); err != nil {
		return "", fmt.Errorf("write object: %w", err)
	}

	return hash, nil
}
