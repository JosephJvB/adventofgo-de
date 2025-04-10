package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Can parse test data", func(t *testing.T) {
		values := readFile("test.txt")
		if len(values) != 10 {
			t.Errorf("Expected 10 values, got %d", len(values))
		}

		if values[0] != 199 {
			t.Errorf("Expected first value to be 199, got %d", values[0])
		}
	})

	t.Run("Can solve test data", func(t *testing.T) {
		values := readFile("test.txt")

		increased := 0
		for i := 1; i < len(values); i++ {
			p := values[i-1]
			c := values[i]

			if c > p {
				increased++
			}
		}

		if increased != 7 {
			t.Errorf("Expected 7 increases, got %d", increased)
		}
	})

	t.Run("Can solve real data", func(t *testing.T) {
		values := readFile("data.txt")

		increased := 0
		for i := 1; i < len(values); i++ {
			p := values[i-1]
			c := values[i]

			if c > p {
				increased++
			}
		}

		fmt.Printf("Part 1: %d\n", increased)

		if increased <= 7 {
			t.Errorf("Expected > 7 increases, got %d", increased)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("test my slicing", func(t *testing.T) {
		values := readFile("test.txt")

		p := values[0:3]
		if len(p) != 3 {
			t.Errorf("Expected 3 values, got %v", p)
		}

		c := values[1:4]
		if len(c) != 3 {
			t.Errorf("Expected 3 values, got %v", c)
		}
	})

	t.Run("Solve test data", func(t *testing.T) {
		values := readFile("test.txt")

		increased := 0
		for i := 3; i < len(values); i++ {
			pl := values[i-3 : i]
			cl := values[i-2 : i+1] // last slice will be length of list
			p := sum(pl)
			c := sum(cl)

			if c > p {
				increased++
			}
		}

		if increased != 5 {
			t.Errorf("Expected 5 increases, got %d", increased)
		}
	})

	t.Run("Solve real data", func(t *testing.T) {
		values := readFile("data.txt")

		increased := 0

		for i := 3; i < len(values); i++ {
			p := sum(values[i-3 : i])
			c := sum(values[i-2 : i+1])

			if c > p {
				increased++
			}
		}

		fmt.Printf("Part 2: %d\n", increased)
		if increased <= 5 {
			t.Errorf("Expected > 5 increases, got %d", increased)
		}
	})
}

func sum(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

func readFile(filename string) []int {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	vals := []int{}
	for i := 0; i < len(lines); i++ {
		s := lines[i]
		if s == "" {
			continue
		}

		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		vals = append(vals, v)
	}

	return vals
}
