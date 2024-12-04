package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(inputFile)
	res := Part1(input)
	fmt.Println("part 1 :", res)

	inputFile, err = os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bufio.NewScanner(inputFile)
	res = Part2(input)
	fmt.Println("part 2 :", res)
}

type direction = struct{ x, y int }

var north = direction{0, -1}
var south = direction{0, 1}
var east = direction{1, 0}
var west = direction{-1, 0}
var northeast = direction{1, -1}
var northwest = direction{-1, -1}
var southeast = direction{1, 1}
var southwest = direction{-1, 1}
var directions = []direction{north, south, east, west, northeast, northwest, southeast, southwest}

func scanDirection(matrix [][]rune, x, y int, dir direction) string {
	res := ""
	for range 3 {
		x += dir.x
		y += dir.y
		if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
			break
		}
		res += string(matrix[y][x])
	}
	return res
}

func Part1(input *bufio.Scanner) int {
	res := 0
	matrix := make([][]rune, 0)
	for input.Scan() {
		row := make([]rune, 0)
		for _, c := range input.Text() {
			row = append(row, c)
		}
		matrix = append(matrix, row)
	}
	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			if matrix[j][i] == 'X' {
				for _, dir := range directions {
					if scanDirection(matrix, i, j, dir) == "MAS" {
						res++
					}
				}
			}
		}
	}

	return res
}

func scanXMAS(matrix [][]rune, x, y int) bool {
	res1 := true
	res1 = res1 && matrix[y+northwest.y][x+northwest.x] == 'M'
	res1 = res1 && matrix[y+southwest.y][x+southwest.x] == 'M'
	res1 = res1 && matrix[y+northeast.y][x+northeast.x] == 'S'
	res1 = res1 && matrix[y+southeast.y][x+southeast.x] == 'S'

	res2 := true
	res2 = res2 && matrix[y+northwest.y][x+northwest.x] == 'M'
	res2 = res2 && matrix[y+southwest.y][x+southwest.x] == 'S'
	res2 = res2 && matrix[y+northeast.y][x+northeast.x] == 'M'
	res2 = res2 && matrix[y+southeast.y][x+southeast.x] == 'S'

	res3 := true
	res3 = res3 && matrix[y+northwest.y][x+northwest.x] == 'S'
	res3 = res3 && matrix[y+southwest.y][x+southwest.x] == 'S'
	res3 = res3 && matrix[y+northeast.y][x+northeast.x] == 'M'
	res3 = res3 && matrix[y+southeast.y][x+southeast.x] == 'M'

	res4 := true
	res4 = res4 && matrix[y+northwest.y][x+northwest.x] == 'S'
	res4 = res4 && matrix[y+southwest.y][x+southwest.x] == 'M'
	res4 = res4 && matrix[y+northeast.y][x+northeast.x] == 'S'
	res4 = res4 && matrix[y+southeast.y][x+southeast.x] == 'M'

	return res1 || res2 || res3 || res4
}

func Part2(input *bufio.Scanner) int {
	res := 0
	matrix := make([][]rune, 0)
	for input.Scan() {
		row := make([]rune, 0)
		for _, c := range input.Text() {
			row = append(row, c)
		}
		matrix = append(matrix, row)
	}
	for j := 1; j < len(matrix)-1; j++ {
		for i := 1; i < len(matrix[j])-1; i++ {
			if matrix[j][i] == 'A' {
				if scanXMAS(matrix, i, j) {
					res++
				}
			}
		}
	}
	return res
}
