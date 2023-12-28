package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/murilo-antas/aoc/day02"
	"github.com/murilo-antas/aoc/util"
)

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

func validateHand(hand *day02.Hand) bool {
    validHand := hand.RedCubes <= maxRedCubes &&
        hand.GreenCubes <= maxGreenCubes &&
        hand.BlueCubes <= maxBlueCubes

    return validHand
}

func validateGame(game *day02.Game) bool {
    for _, hand := range game.Hands {
        if !validateHand(&hand) {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("../input.txt")
    util.Check(err)
    defer file.Close()

    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        game := day02.ParseGame(line)
        validGame := validateGame(game)
        if validGame {
            sum += game.Id
        }
    }

    err = scanner.Err()
    util.Check(err)

    fmt.Println(sum)
}
