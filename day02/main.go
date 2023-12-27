package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/murilo-antas/aoc/util"
)

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

type game struct {
    id int
    hands []hand
}

type hand struct {
    redCubes int
    greenCubes int
    blueCubes int
}

func parseHand(handString string) *hand {
    var hand hand
    cubes := strings.Split(handString, ",")
    regexp := regexp.MustCompile(`(\d+) (.+)`)
    for _, cube := range cubes {
        res := regexp.FindAllStringSubmatch(strings.TrimSpace(cube), -1)
        if res == nil {
            panic("Invalid hand")
        }

        for i := range res {
            cubeValue, err := strconv.Atoi(res[i][1])
            util.Check(err)
            switch res[i][2] {
            case "red":
                hand.redCubes = cubeValue
            case "green":
                hand.greenCubes = cubeValue
            case "blue":
                hand.blueCubes = cubeValue
            }
        }
    }
    return &hand
}

func parseHands(handsString string) []hand {
    hs := make([]hand, 1)
    hands := strings.Split(handsString, ";") 
    for _, hand := range hands {
        h := parseHand(hand)
        hs = append(hs, *h)
    }
    return hs
}

func parseGame(line string) *game {
    var game game
    var err error

    regexp := regexp.MustCompile(`Game (\d+): (.+)`)
    res := regexp.FindAllStringSubmatch(line, -1)
    if res == nil {
        panic("Invalid format")
    }

    for i := range res {
        game.id, err = strconv.Atoi(res[i][1])
        util.Check(err)

        handsString := res[i][2]
        game.hands = parseHands(handsString)
    }

    return &game
}

func validateHand(hand *hand) bool {
    validHand := hand.redCubes <= maxRedCubes &&
        hand.greenCubes <= maxGreenCubes &&
        hand.blueCubes <= maxBlueCubes

    return validHand
}

func validateGame(game *game) bool {
    for _, hand := range game.hands {
        if !validateHand(&hand) {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("./input.txt")
    util.Check(err)
    defer file.Close()

    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        game := parseGame(line)
        validGame := validateGame(game)
        if validGame {
            sum += game.id
        }
    }

    err = scanner.Err()
    util.Check(err)

    fmt.Println(sum)
}
