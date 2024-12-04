package main

import (
	"bufio"
	"strings"
	"testing"
)

const example string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const expectedResultPart1 int = 18
const expectedResultPart2 int = 9

func TestPart1(t *testing.T) {
	input := strings.NewReader(example)
	res := Part1(bufio.NewScanner(input))
	if res != expectedResultPart1 {
		t.Errorf("got %d, expected %d", res, expectedResultPart1)
	}
}

func TestPart2(t *testing.T) {
	input := strings.NewReader(example)
	res := Part2(bufio.NewScanner(input))
	if res != expectedResultPart2 {
		t.Errorf("got %d, expected %d", res, expectedResultPart2)
	}
}
