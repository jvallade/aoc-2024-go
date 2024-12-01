package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	// inputFile, err = os.Open("day1/input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// input = bufio.NewScanner(inputFile)
	// res = Part2(input)
	// fmt.Println("part 2 :", res)
}

func Part1(input *bufio.Scanner) int {
	for input.Scan() {
	}

	return 0
}

func Part2(input *bufio.Scanner) int {
	res := 0
	return res
}
