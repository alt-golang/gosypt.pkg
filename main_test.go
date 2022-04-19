package gosypt_pkg

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.65 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestEncryptDecrypt(t *testing.T) {

	encrytped, _ := EncryptString("1234567890123456", "HelloWorld")
	decrytped, _ := DecryptString("1234567890123456", encrytped)

	if decrytped != "HelloWorld" {
		t.Errorf("decrytped != \"HelloWorld\": decrytped is:%s", decrytped)
	}

	_, err := EncryptString("123456789012345", "HelloWorld")

	if err == nil {
		t.Errorf("err == nil: err is:%s", "nil")
	}
}
