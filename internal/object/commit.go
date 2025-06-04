package object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func WriteCommit(message, author string, parent *string) (string, error) {
	treeSHA, err := WriteTree(".")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "tree %s\n", treeSHA)
  if parent != nil {
      fmt.Fprintf(&buf, "parent %s\n", *parent)
  }

  ts := time.Now()
  fmt.Fprintf(&buf, "author %s %d +0530\n", author, ts.Unix())
  fmt.Fprintf(&buf, "committer %s %d +0530\n", author, ts.Unix())
  buf.WriteByte('\n')
  buf.WriteString(message)
  content := buf.Bytes()

  header := fmt.Sprintf("commit %d\x00", len(content))
  full := append([]byte(header), content...)

  sha := fmt.Sprintf("%x", sha1.Sum(full))
  dir := filepath.Join(".git", "objects", sha[:2])
  file := filepath.Join(dir, sha[2:])

  os.MkdirAll(dir, 0755)

  var compressed bytes.Buffer
  w := zlib.NewWriter(&compressed)
  w.Write(full)
  w.Close()

  os.WriteFile(file, compressed.Bytes(), 0644)

  return sha, nil
}
