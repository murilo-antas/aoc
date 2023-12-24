package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1() {
    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        line := []rune(scanner.Text())
        var digit1 rune 
        var digit2 rune 
        for _, r := range line {
            if unicode.IsDigit(r) {
                digit1 = r
                break
            }
        }
        for i := len(line) - 1; i >= 0; i-- {
            if unicode.IsDigit(line[i]) {
                digit2 = line[i]
                break
            }
        }
        s := string(digit1) + string(digit2)
        n, err :=  strconv.Atoi(s)
        check(err)
        sum += n
    }
    err = scanner.Err()
    check(err)

    fmt.Println(sum)
}

func part2() {

    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        line := []rune(scanner.Text())
        var digit1 int 
        var digit2 int 
        var tokens []int
        for i, r := range line {
            if unicode.IsDigit(r) {
                d, _ := strconv.Atoi(string(r))
                tokens = append(tokens, d)
            } else {
                d, found := findDigitString(line[i:])
                if found {
                    tokens = append(tokens, d)
                }
            }
        }
        digit1 = tokens[0]
        digit2 = tokens[len(tokens) - 1]
        s := fmt.Sprintf("%d%d", digit1, digit2)
        n, err :=  strconv.Atoi(s)
        check(err)
        sum += n
    }
    err = scanner.Err()
    check(err)

    fmt.Println(sum)
}

func findDigitString(s []rune) (int, bool) {
    digits := map[string]int {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    for k, v := range digits {
        if strings.HasPrefix(string(s), k) {
            return v, true
        }
    }
    return 0, false 
}

func main() {
    part1()
    part2()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
