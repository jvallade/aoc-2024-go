package main

import (
	"bufio"
	"strings"
	"testing"
)

const example string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

const expectedResultPart1 int = 143
const expectedResultPart2 int = 123

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
