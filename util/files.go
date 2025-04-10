package util

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	linesTrim := []string{}
	for i := 0; i < len(lines); i++ {
		s := lines[i]
		if s == "" {
			continue
		}

		linesTrim = append(linesTrim, s)
	}

	return linesTrim
}
