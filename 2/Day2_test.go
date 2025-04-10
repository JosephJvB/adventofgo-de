package main

import (
	"aog-2021/util"
	"log"
	"strconv"
	"strings"
	"testing"
)

type Move struct {
	direction string
	distance  int
}

type Coord struct {
	x int
	y int
}

type Sub struct {
	Coord
	aim int
}

func TestPart1(t *testing.T) {
	t.Run("Can parse test data", func(t *testing.T) {
		moves := parseMoves("test.txt")

		if len(moves) != 6 {
			t.Errorf("Expected 6 moves, got %d", len(moves))
		}

		if moves[0].direction != "forward" || moves[0].distance != 5 {
			t.Errorf("Expected first move to be forward 5, got %s %d", moves[0].direction, moves[0].distance)
		}

		if moves[5].direction != "forward" || moves[5].distance != 2 {
			t.Errorf("Expected first move to be forward 2, got %s %d", moves[0].direction, moves[0].distance)
		}
	})

	t.Run("Can solve test data", func(t *testing.T) {
		moves := parseMoves("test.txt")

		pos := Coord{x: 0, y: 0}

		for _, move := range moves {
			pos = updateCoord(pos, move)
		}

		res := pos.x * pos.y
		if res != 150 {
			t.Errorf("Expected 150, got %d from coord %v", res, pos)
		}
	})

	t.Run("Can solve part 1", func(t *testing.T) {
		moves := parseMoves("data.txt")

		pos := Coord{x: 0, y: 0}

		for _, move := range moves {
			pos = updateCoord(pos, move)
		}

		res := pos.x * pos.y
		t.Logf("Part 1: %d", res)

		if res != 1868935 {
			t.Errorf("Expected 1868935, got %d from coord %v", res, pos)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("Can solve test data", func(t *testing.T) {
		moves := parseMoves("test.txt")

		sub := Sub{
			Coord: Coord{x: 0, y: 0},
			aim:   0,
		}

		for _, move := range moves {
			sub = updateSubCoord(sub, move)
		}

		res := sub.x * sub.y
		if res != 900 {
			t.Errorf("Expected 900, got %d from sub %v", res, sub)
		}
	})

	t.Run("Can solve part 2 data", func(t *testing.T) {
		moves := parseMoves("data.txt")

		sub := Sub{
			Coord: Coord{x: 0, y: 0},
			aim:   0,
		}

		for _, move := range moves {
			sub = updateSubCoord(sub, move)
		}

		res := sub.x * sub.y
		log.Printf("Part 2: %d", res)

		if res != 1965970888 {
			t.Errorf("Expected 1965970888, got %d from sub %v", res, sub)
		}
	})
}

func updateCoord(coord Coord, move Move) Coord {
	switch move.direction {
	case "forward":
		coord.x += move.distance
	case "down":
		coord.y += move.distance
	case "up":
		coord.y -= move.distance
	default:
		panic("Unknown direction")
	}

	return coord
}

func updateSubCoord(sub Sub, move Move) Sub {
	switch move.direction {
	case "forward":
		sub.x += move.distance
		sub.y += (sub.aim * move.distance)
	case "down":
		sub.aim += move.distance
	case "up":
		sub.aim -= move.distance
	default:
		panic("Unknown direction")
	}

	return sub
}

func parseMoves(filename string) []Move {
	lines := util.ReadFile(filename)

	moves := []Move{}
	for _, line := range lines {
		fields := strings.Fields(line)

		distance, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		m := Move{
			direction: fields[0],
			distance:  distance,
		}

		moves = append(moves, m)
	}

	return moves
}
