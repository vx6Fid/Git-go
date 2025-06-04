package object

import (
	"os"
	"testing"
	"path/filepath"
)

func TestHashAndStoreBlob(t *testing.T) {
	tmp := t.TempDir()
	os.Chdir(tmp)

	os.MkdirAll(".git/objects", 0755)
	testFile := "hello.txt"
	os.WriteFile(testFile, []byte("hello world\n"), 0644)

	hash, err := HashAndStoreBlob(testFile)
	if err != nil {
		t.Fatalf("HashAndStoreBlob failed: %v", err)
	}
	
	dir := filepath.Join(".git", "objects", hash[:2])
	file := filepath.Join(dir, hash[2:])
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("Expected object file %s to exist", file)
	}
}

func BenchmarkHashAndStoreBlob(b *testing.B) {
    content := []byte("hello world\n")
    file := "bench.txt"

    os.WriteFile(file, content, 0644)
    os.MkdirAll(".git/objects", 0755)

    b.ResetTimer() 

    for i := 0; i < b.N; i++ {
        HashAndStoreBlob(file)
    }
}

