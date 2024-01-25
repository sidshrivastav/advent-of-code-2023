package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

func isConfigurationPossible(game string) int {
	red := 0
	green := 0
	blue := 0
	gameId := strings.Split(strings.Split(game, ":")[0], " ")[1]
	gameData := strings.Split(strings.Split(game, ":")[1], ";")
	for _, gameSet :=  range gameData {
		for _, gameCubes := range strings.Split(gameSet, ",") {
			gameCube := strings.Split(gameCubes, " ")
			cubeCount, err := strconv.Atoi(gameCube[1])
			if err != nil {
				panic(err)
			}
			if gameCube[2] == "red" {
				red += cubeCount
			}
			if gameCube[2] == "green" {
				green += cubeCount
			}
			if gameCube[2] == "blue" {
				blue += cubeCount
			}
		}
	}

	if red <= 12 && green <= 13 && blue <= 14 {
		gameIdNumber, err := strconv.Atoi(gameId)
		if err != nil {
			panic(err)
		}
		return gameIdNumber
	}

	return 0
}

func solve(part string, gameRound []string) (result int) {
	for _, value := range gameRound {
		if part == "1" && value != "" {
			result += isConfigurationPossible(value)
		}
	}

	return
}

func main() {
	// Part 1
	inputFile := "input_1.txt"
	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	gameRound := strings.Split(string(content), "\n")
	fmt.Println(solve("1", gameRound))
}