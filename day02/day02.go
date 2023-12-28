package day02

import (
    "strings"
    "regexp"
    "strconv"
    "github.com/murilo-antas/aoc/util"
)

type Game struct {
    Id int
    Hands []Hand
}

type Hand struct {
    RedCubes int
    GreenCubes int
    BlueCubes int
}


func ParseHand(handString string) *Hand {
    var hand Hand

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
                hand.RedCubes = cubeValue
            case "green":
                hand.GreenCubes = cubeValue
            case "blue":
                hand.BlueCubes = cubeValue
            }
        }
    }
    return &hand
}

func ParseHands(handsString string) []Hand {
    hs := make([]Hand, 1)
    hands := strings.Split(handsString, ";") 
    for _, hand := range hands {
        h := ParseHand(hand)
        hs = append(hs, *h)
    }
    return hs
}

func ParseGame(line string) *Game {
    var game Game
    var err error

    regexp := regexp.MustCompile(`Game (\d+): (.+)`)
    res := regexp.FindAllStringSubmatch(line, -1)
    if res == nil {
        panic("Invalid format")
    }

    for i := range res {
        game.Id, err = strconv.Atoi(res[i][1])
        util.Check(err)

        handsString := res[i][2]
        game.Hands = ParseHands(handsString)
    }

    return &game
}
