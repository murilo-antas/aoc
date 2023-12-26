package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDigits(line []rune, validDigits map[string]int) (digit1, digit2 int) {
	var tokens []int

	for i := range line {
		d, found := findDigitString(line[i:], validDigits)
		if found {
			tokens = append(tokens, d)
		}
	}
	digit1 = tokens[0]
	digit2 = tokens[len(tokens)-1]
	return digit1, digit2
}

func findDigitString(line []rune, validDigits map[string]int) (int, bool) {
	for k, v := range validDigits {
		if strings.HasPrefix(string(line), k) {
			return v, true
		}
	}
	return 0, false
}

func getSum(validDigits map[string]int) {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		digit1, digit2 := getDigits(line, validDigits)
		s := fmt.Sprintf("%d%d", digit1, digit2)
		n, err := strconv.Atoi(s)
		check(err)
		sum += n
	}
	err = scanner.Err()
	check(err)

	fmt.Println(sum)
}

var validDigitsPart1 = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"0": 0,
}

var validDigitsPart2 = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	getSum(validDigitsPart1)
	getSum(validDigitsPart2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
