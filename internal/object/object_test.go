// object_test.go
package object

import (
    "os"
    "testing"
		"bytes"
)

func TestWriteAndReadBlob(t *testing.T) {
    content := []byte("hello test blob\n")
    tmpFile := "temp.txt"
    os.WriteFile(tmpFile, content, 0644)
    defer os.Remove(tmpFile)

    sha, err := HashAndStoreBlob(tmpFile)
    if err != nil {
        t.Fatal(err)
    }

    typ, data, err := ReadObject(sha)
    if err != nil {
        t.Fatal(err)
    }

    if typ != "blob" {
        t.Errorf("Expected blob, got %s", typ)
    }

    if string(data) != string(content) {
        t.Errorf("Blob content mismatch: got %q", data)
    }
}

func TestWriteAndReadTree(t *testing.T) {
    os.Mkdir("testdir", 0755)
    os.WriteFile("testdir/file1.txt", []byte("foo"), 0644)
    os.WriteFile("testdir/file2.txt", []byte("bar"), 0644)
    defer os.RemoveAll("testdir")

    // Save objects
    sha, err := WriteTree("testdir")
    if err != nil {
        t.Fatal(err)
    }

    os.RemoveAll("testdir") // simulate deleting working dir

    err = ReadTree(sha, "testdir")
    if err != nil {
        t.Fatal(err)
    }

    // Verify files exist
    checkFile := func(path, expected string) {
        data, err := os.ReadFile(path)
        if err != nil {
            t.Fatalf("missing file: %s", path)
        }
        if string(data) != expected {
            t.Errorf("file %s has wrong content: %q", path, data)
        }
    }

    checkFile("testdir/file1.txt", "foo")
    checkFile("testdir/file2.txt", "bar")
}

func TestWriteCommit(t *testing.T) {
    os.WriteFile("a.txt", []byte("test commit"), 0644)
    defer os.Remove("a.txt")

    _, err := HashAndStoreBlob("a.txt")
    if err != nil {
        t.Fatal(err)
    }

    treeSha, err := WriteTree(".")
    if err != nil {
        t.Fatal(err)
    }

    sha, err := WriteCommit("msg", "Achal <me@example.com>", nil)
    if err != nil {
        t.Fatal(err)
    }

    typ, content, err := ReadObject(sha)
    if err != nil || typ != "commit" {
        t.Fatalf("Invalid commit read: %s", err)
    }

    if !bytes.Contains(content, []byte(treeSha)) {
        t.Errorf("Commit doesn't reference expected tree SHA")
    }
}

