package main

import (
	"aog-2021/util"
	"bytes"
	"strconv"
	"testing"
)

type Power struct {
	gamma string
	eps   string
}

func getPowerBinaries(lines []string) Power {
	var gamma bytes.Buffer
	var eps bytes.Buffer

	half := len(lines) / 2

	for i := 0; i < len(lines[0]); i++ {
		ones := 0
		for _, line := range lines {
			c := line[i]

			// use single quotes for "byte". Like "char" from C#
			if c == '1' {
				ones++
			}
		}

		// fmt.Printf("col:%d ones:%d\n", i, ones)

		if ones > half {
			gamma.WriteByte('1')
			eps.WriteByte('0')
		} else {
			gamma.WriteByte('0')
			eps.WriteByte('1')
		}
	}

	return Power{
		gamma: gamma.String(),
		eps:   eps.String(),
	}
}

func TestPart1(t *testing.T) {
	t.Run("Can parse test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		if lines[0] != "00100" {
			t.Errorf("Expected first line to be 00100, got %s", lines[0])
		}

		if len(lines) != 12 {
			t.Errorf("Expected 12 lines, got %d", len(lines))
		}
	})

	t.Run("getPowerBinaries works for 00100", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		power := getPowerBinaries(lines)

		if power.gamma != "10110" {
			t.Errorf("expected gamma to be 10110 got %s", power.gamma)
		}
		if power.eps != "01001" {
			t.Errorf("expected eps to be 01001 got %s", power.eps)
		}
	})

	t.Run("Can solve test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		powerBinaries := getPowerBinaries(lines)

		gamma, err := strconv.ParseInt(powerBinaries.gamma, 2, 64)
		if err != nil {
			panic(err)
		}
		if gamma != 22 {
			t.Errorf("expected gamma to be 22 got %d", gamma)
		}

		eps, err := strconv.ParseInt(powerBinaries.eps, 2, 64)
		if err != nil {
			panic(err)
		}
		if eps != 9 {
			t.Errorf("expected eps to be 9 got %d", eps)
		}

		power := gamma * eps
		if power != 198 {
			t.Errorf("Expected power to be 198 got %d", power)
		}
	})

	t.Run("Can solve Part 1", func(t *testing.T) {
		lines := util.ReadFile("data.txt")

		powerBinaries := getPowerBinaries(lines)

		gamma, err := strconv.ParseInt(powerBinaries.gamma, 2, 64)
		if err != nil {
			panic(err)
		}

		eps, err := strconv.ParseInt(powerBinaries.eps, 2, 64)
		if err != nil {
			panic(err)
		}

		power := gamma * eps

		t.Logf("Part 1:%d", power)

		if power != 2583164 {
			t.Errorf("expected power to be 2583164 got %d", power)
		}
	})
}

func getOx(lines []string) string {
	ox := lines

	for i := 0; i < len(lines[0]); i++ {
		if len(ox) == 1 {
			break
		}

		half := len(ox) / 2
		ones := 0
		zeros := 0
		for _, line := range ox {
			if ones > half {
				break
			}
			if zeros > half {
				break
			}

			if line[i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		isOnes := ones > half || ones == zeros
		next := []string{}
		for _, line := range ox {
			c := line[i]

			if isOnes && c == '1' {
				next = append(next, line)
			}
			if !isOnes && c == '0' {
				next = append(next, line)
			}
		}

		ox = next
	}

	return ox[0]
}

func getCo2(lines []string) string {
	co2 := lines

	for i := 0; i < len(lines[0]); i++ {
		if len(co2) == 1 {
			break
		}

		half := len(co2) / 2
		ones := 0
		zeros := 0
		for _, line := range co2 {
			if ones > half {
				break
			}
			if zeros > half {
				break
			}

			if line[i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		isOnes := ones > half || ones == zeros
		next := []string{}
		for _, line := range co2 {
			c := line[i]

			if isOnes && c == '0' {
				next = append(next, line)
			}
			if !isOnes && c == '1' {
				next = append(next, line)
			}
		}

		co2 = next
	}

	return co2[0]
}

func TestPart2(t *testing.T) {
	t.Run("can get ox from test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		ox := getOx(lines)

		if ox != "10111" {
			t.Errorf("Expected ox to be 10111. Got %s", ox)
		}
	})

	t.Run("can get co2 from test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		co2 := getCo2(lines)

		if co2 != "01010" {
			t.Errorf("Expected co2 to be 01010. Got %s", co2)
		}
	})

	t.Run("can solve test data", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		ox := getOx(lines)
		co2 := getCo2(lines)

		oxV, err := strconv.ParseInt(ox, 2, 64)
		if err != nil {
			panic(err)
		}
		co2V, err := strconv.ParseInt(co2, 2, 64)
		if err != nil {
			panic(err)
		}

		v := oxV * co2V
		if v != 230 {
			t.Errorf("Expected 230. Got %d", v)
		}
	})

	t.Run("can solve real data", func(t *testing.T) {
		lines := util.ReadFile("data.txt")

		ox := getOx(lines)
		co2 := getCo2(lines)

		oxV, err := strconv.ParseInt(ox, 2, 64)
		if err != nil {
			panic(err)
		}
		co2V, err := strconv.ParseInt(co2, 2, 64)
		if err != nil {
			panic(err)
		}

		v := oxV * co2V
		if v != 2784375 {
			t.Errorf("Expected 2784375. Got %d", v)
		}
	})
}

func TestPart2Refactor(t *testing.T) {
	t.Run("Can refactor my solution", func(t *testing.T) {
		lines := util.ReadFile("test.txt")

		lifeSupport := getLifeSupport(lines)
		ox := lifeSupport[0]
		if ox != "10111" {
			t.Errorf("Expected 10111 got %s", ox)
		}

		co2 := lifeSupport[1]
		if co2 != "01010" {
			t.Errorf("Expected 01010 got %s", co2)
		}
	})

	t.Run("refactored can solve real data", func(t *testing.T) {
		lines := util.ReadFile("data.txt")

		lifeSupport := getLifeSupport(lines)
		ox := lifeSupport[0]
		co2 := lifeSupport[1]

		oxV, err := strconv.ParseInt(ox, 2, 64)
		if err != nil {
			panic(err)
		}
		co2V, err := strconv.ParseInt(co2, 2, 64)
		if err != nil {
			panic(err)
		}

		v := oxV * co2V
		if v != 2784375 {
			t.Errorf("Expected 2784375. Got %d", v)
		}
	})
}

func getLifeSupport(lines []string) []string {
	ox := lines
	co2 := lines

	for i := 0; i < len(lines[0]); i++ {
		lox := len(ox)
		lco2 := len(co2)
		if lox == 1 && lco2 == 1 {
			break
		}

		if lox > 1 {
			ox = filterLines(i, ox, func(zeros int, ones int, b byte) bool {
				return keepOxygen(zeros, ones, b)
			})
		}

		if lco2 > 1 {
			co2 = filterLines(i, co2, func(zeros int, ones int, b byte) bool {
				return keepCo2(zeros, ones, b)
			})
		}
	}

	return []string{ox[0], co2[0]}
}

func keepOxygen(zeros int, ones int, b byte) bool {
	if b == '1' {
		return ones > zeros || ones == zeros
	} else {
		return zeros > ones
	}
}
func keepCo2(zeros int, ones int, b byte) bool {
	if b == '0' {
		return ones > zeros || ones == zeros
	} else {
		return zeros > ones
	}
}

func filterLines(i int, lines []string, cb func(zeros int, ones int, c byte) bool) []string {
	zeros := 0
	ones := 0
	for _, line := range lines {
		if line[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}

	next := []string{}
	for _, line := range lines {
		if cb(zeros, ones, line[i]) {
			next = append(next, line)
		}
	}

	return next
}
