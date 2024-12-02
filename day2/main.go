package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day2/input.txt")
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
		levels := make([]int, 0)
		for _, l := range strings.Fields(input.Text()) {
			intLevel, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, intLevel)
		}
		isSafe := true
		isIncreasing := false
		isDecreasing := false
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]
			if diff == 0 || diff > 3 || diff < -3 {
				isSafe = false
				break
			}
			if diff > 0 {
				isIncreasing = true
			}
			if diff < 0 {
				isDecreasing = true
			}
			if isIncreasing && isDecreasing {
				isSafe = false
				break
			}
		}
		if isSafe {
			res++
		}

	}

	return res
}

func isSequenceSafe(levels []int) bool {
	isIncreasing := false
	isDecreasing := false
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if diff > 0 {
			isIncreasing = true
		}
		if diff < 0 {
			isDecreasing = true
		}
		if isIncreasing && isDecreasing {
			return false
		}
	}
	return true
}

func Part2(input *bufio.Scanner) int {
	res := 0
	for input.Scan() {
		levels := make([]int, 0)
		for _, l := range strings.Fields(input.Text()) {
			intLevel, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, intLevel)
		}
		isSafe := isSequenceSafe(levels)
		if !isSafe {
			for i := 0; i < len(levels); i++ {
				newLevels := slices.Clone(levels)
				newLevels = slices.Delete(newLevels, i, i+1)
				isSafe = isSequenceSafe(newLevels)
				if isSafe {
					break
				}
			}
		}

		if isSafe {
			res++
		}
	}

	return res
}
