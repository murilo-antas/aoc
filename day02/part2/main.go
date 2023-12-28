package main 

import (
	"bufio"
	"fmt"
	"os"

	"github.com/murilo-antas/aoc/day02"
	"github.com/murilo-antas/aoc/util"
)

func getGameScore(game *day02.Game) int {
    minRed := 0
    minGreen := 0
    minBlue := 0
    for _, hand := range game.Hands {
        if minRed < hand.RedCubes {
            minRed = hand.RedCubes
        }
        if minGreen < hand.GreenCubes {
            minGreen = hand.GreenCubes
        }
        if minBlue < hand.BlueCubes {
            minBlue = hand.BlueCubes
        }
    }
    return minRed * minGreen * minBlue
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
                gameScore := getGameScore(game)
		sum += gameScore
	}
	err = scanner.Err()
	util.Check(err)

	fmt.Println(sum)
}
