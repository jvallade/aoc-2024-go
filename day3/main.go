package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

func Part1(input *bufio.Scanner) int {
	res := 0
	for input.Scan() {
		line := input.Text()
		re, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			val1, _ := strconv.Atoi(match[1])
			val2, _ := strconv.Atoi(match[2])
			res += val1 * val2
		}
	}
	return res
}

func Part2(input *bufio.Scanner) int {
	res := 0
	oneLineInput := ""
	for input.Scan() {
		oneLineInput += input.Text()
	}

	oneLineInput += "do()"
	reFilter, _ := regexp.Compile(`don't\(\).+?do\(\)`)
	oneLineInput = reFilter.ReplaceAllString(oneLineInput, "")

	re, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(oneLineInput, -1)
	for _, match := range matches {
		val1, _ := strconv.Atoi(match[1])
		val2, _ := strconv.Atoi(match[2])
		res += val1 * val2
	}

	return res
}
