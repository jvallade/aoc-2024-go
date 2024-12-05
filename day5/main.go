package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	inputFile, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

type PageOrder struct {
	before int
	after  int
}

type Update []int

func (u Update) isOrdered(orders []PageOrder) bool {
	for _, o := range orders {
		idxBefore := slices.Index(u, o.before)
		idxAfter := slices.Index(u, o.after)
		if idxBefore == -1 || idxAfter == -1 {
			continue
		}
		if idxBefore > idxAfter {
			return false
		}
	}
	return true
}

func (u *Update) sort(orders []PageOrder) {
	for range len(*u) {
		for _, o := range orders {
			idxBefore := slices.Index(*u, o.before)
			idxAfter := slices.Index(*u, o.after)
			if idxBefore == -1 || idxAfter == -1 {
				continue
			}
			if idxBefore > idxAfter {
				(*u)[idxBefore], (*u)[idxAfter] = (*u)[idxAfter], (*u)[idxBefore]
			}
		}
		if u.isOrdered(orders) {
			break
		}
	}
}

func parseInput(input *bufio.Scanner) ([]PageOrder, []Update) {
	var orders []PageOrder
	var updates []Update
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		var order PageOrder
		fmt.Sscanf(line, "%d|%d", &order.before, &order.after)
		orders = append(orders, order)
	}
	for input.Scan() {
		line := input.Text()
		var update Update
		for _, p := range strings.Split(line, ",") {
			var page int
			fmt.Sscanf(p, "%d", &page)
			update = append(update, page)
		}
		updates = append(updates, update)
	}
	return orders, updates
}

func Part1(input *bufio.Scanner) int {
	res := 0
	orders, updates := parseInput(input)
	for _, u := range updates {
		if u.isOrdered(orders) {
			res += u[(len(u)-1)/2]
		}
	}
	return res
}

func Part2(input *bufio.Scanner) int {
	res := 0
	orders, updates := parseInput(input)
	for _, u := range updates {
		if !u.isOrdered(orders) {
			u.sort(orders)
			res += u[(len(u)-1)/2]
		}
	}
	return res
}
