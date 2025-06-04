package object 
import (
	"testing"
	"os"
)

func TestReadObject(t *testing.T) {
	tmp := t.TempDir()
	os.Chdir(tmp)

	os.MkdirAll(".git/objects", 0755)
	os.WriteFile("test.txt", []byte("hello world\n"), 0644)

	hash, err := HashAndStoreBlob("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	objType, content, err := ReadObject(hash)
	if err != nil {
		t.Fatal(err)
	}

	if objType != "blob" {
		t.Errorf("Expected blob, got %s", objType)
	}

	expected := "hello world\n"
	if string(content) != expected {
		t.Errorf("Expected content %q, got %q", expected, string(content))
	}

}
