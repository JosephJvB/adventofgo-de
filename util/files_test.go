package util

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("ReadFile", func(t *testing.T) {
		lines := ReadFile("../1/test.txt")

		if len(lines) != 10 {
			t.Errorf("Expected 10 lines, got %d", len(lines))
		}

		if lines[0] != "199" {
			t.Errorf("Expected first value to be 199, got %s", lines[0])
		}
	})

	//////
	// works but only cos there's no spaces in the lines
	/////
	t.Run("strings Fields 1/test.txt", func(t *testing.T) {
		bytes, err := os.ReadFile("../1/test.txt")
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Fields(string(bytes))

		if len(lines) != 10 {
			t.Errorf("Expected 10 lines, got %d", len(lines))
		}
	})

	//////
	// Fails
	/////
	t.Run("strings Fields 2/test.txt", func(t *testing.T) {
		t.Skip("This test proves an invalid assumption")

		bytes, err := os.ReadFile("../2/test.txt")
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Fields(string(bytes))

		if len(lines) != 6 {
			t.Errorf("Expected 6 lines, got %d %v", len(lines), lines)
		}
	})

}
