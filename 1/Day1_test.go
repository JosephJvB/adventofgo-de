package main

import (
	"aog-2021/util"
	"strconv"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	t.Run("Can parse test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

		if len(values) != 10 {
			t.Errorf("Expected 10 values, got %d", len(values))
		}

		if values[0] != 199 {
			t.Errorf("Expected first value to be 199, got %d", values[0])
		}
	})

	t.Run("Can solve test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

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
		lines := util.ReadFile("data.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

		increased := 0
		for i := 1; i < len(values); i++ {
			p := values[i-1]
			c := values[i]

			if c > p {
				increased++
			}
		}

		t.Logf("Part 1: %d\n", increased)

		if increased <= 7 {
			t.Errorf("Expected > 7 increases, got %d", increased)
		}
	})
}

func TestDay1Part2(t *testing.T) {
	t.Run("test my slicing", func(t *testing.T) {
		lines := util.ReadFile("test.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

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
		lines := util.ReadFile("test.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

		increased := 0
		for i := 3; i < len(values); i++ {
			pl := values[i-3 : i]
			cl := values[i-2 : i+1] // last slice will be length of list
			p := util.Sum(pl)
			c := util.Sum(cl)

			if c > p {
				increased++
			}
		}

		if increased != 5 {
			t.Errorf("Expected 5 increases, got %d", increased)
		}
	})

	t.Run("Solve real data", func(t *testing.T) {
		lines := util.ReadFile("data.txt")
		values := []int{}
		for _, line := range lines {
			s, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			values = append(values, s)
		}

		increased := 0

		for i := 3; i < len(values); i++ {
			p := util.Sum(values[i-3 : i])
			c := util.Sum(values[i-2 : i+1])

			if c > p {
				increased++
			}
		}

		t.Logf("Part 2: %d\n", increased)
		if increased <= 5 {
			t.Errorf("Expected > 5 increases, got %d", increased)
		}
	})
}
