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
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for input.Scan() {
		fields := strings.Fields(input.Text())
		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	res := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] >= list2[i] {
			res += list1[i] - list2[i]
		} else {
			res += list2[i] - list1[i]
		}
	}
	return res
}

func Part2(input *bufio.Scanner) int {
	res := 0
	return res
}
