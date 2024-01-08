package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/murilo-antas/aoc/util"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

var Directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func CheckSymbol(c byte) bool {
	return !unicode.IsDigit(rune(c)) && rune(c) != '.' && rune(c) != '\n'
}

func CheckAdjacentSymbol(content []string, x int, y int) bool {
	for _, d := range Directions {
		dx, dy := d[0], d[1]
		if y+dy < 0 || y+dy >= len(content) || x+dx < 0 || x+dx >= len(content[y+dy]) {
			continue
		}
		if CheckSymbol(content[y+dy][x+dx]) {
			return true
		}
	}
	return false
}

func main() {
	content, err := ReadInput("../input.txt")
	util.Check(err)

	sum := 0
	var n int
	for y, line := range content {
		currentNumber := make([]byte, 0)
		hasAdjacentSymbol := false
		for x := range line {
			if unicode.IsDigit(rune(line[x])) {
				currentNumber = append(currentNumber, line[x])
				if !hasAdjacentSymbol {
					hasAdjacentSymbol = CheckAdjacentSymbol(content, x, y)
				}
			} 

			if (!unicode.IsDigit(rune(line[x])) || x == len(line) - 1) && len(currentNumber) > 0 && hasAdjacentSymbol {
				n, err = strconv.Atoi(strings.TrimSpace(string(currentNumber)))
				util.Check(err)
                                fmt.Println(n)
				sum += n
                                currentNumber = currentNumber[:0]
				hasAdjacentSymbol = false
			}

                        if !unicode.IsDigit(rune(line[x])) {
                            currentNumber = currentNumber[:0]
                        }
		}
	}
	fmt.Println(sum)
}
