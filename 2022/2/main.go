package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cleanedSlice := readInputFile("input-1.txt")
	rawPoints := calculatePointsV2(cleanedSlice)
	output := aggregatePoints(rawPoints)
	fmt.Println(output)
}

func readInputFile(fileName string) [][]string {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	pairedStrings := strings.Split(input, "\n")

	sliceLength := len(pairedStrings)

	plays := make([][]string, sliceLength)

	/*
		Creates matrix of plays
	*/
	for i, val := range pairedStrings {
		playSlice := strings.Split(val, " ")
		plays[i] = playSlice
	}

	return plays
}

func calculatePoints(plays [][]string) []int64 {
	sliceLength := len(plays)
	pointsOutput := make([]int64, sliceLength)

	for i, val := range plays {
		fmt.Println(val)
		opponentPlay := val[0]
		myPlay := val[1]
		basePoints := calculateBasePoints(myPlay)
		playPoints := calculatePlayPoints(myPlay, opponentPlay)
		pointsOutput[i] = basePoints + playPoints
	}

	return pointsOutput
}

func calculateBasePoints(myPlay string) int64 {
	var points int64 = 0

	switch myPlay {
	case "X":
		points = 1
	case "Y":
		points = 2
	case "Z":
		points = 3
	}

	fmt.Printf("calculateBasePoints: %d\n", points)

	return points
}

func calculatePlayPoints(myPlay, opponentPlay string) int64 {
	var points int64 = 0
	switch {
	case myPlay == "X" && opponentPlay == "C":
		points = 6
	case myPlay == "Y" && opponentPlay == "A":
		points = 6
	case myPlay == "Z" && opponentPlay == "B":
		points = 6
	case myPlay == "X" && opponentPlay == "A":
		points = 3
	case myPlay == "Y" && opponentPlay == "B":
		points = 3
	case myPlay == "Z" && opponentPlay == "C":
		points = 3
	default:
		points = 0
	}

	fmt.Printf("calculatePlayPoints: %d\n\n", points)

	return points
}

func aggregatePoints(rawPoints []int64) int64 {
	var totalPoints int64 = 0
	for _, val := range rawPoints {
		totalPoints += val
	}
	return totalPoints
}

func calculatePointsV2(plays [][]string) []int64 {
	sliceLength := len(plays)
	pointsOutput := make([]int64, sliceLength)

	for i, val := range plays {
		fmt.Println(val)
		opponentPlay := val[0]
		myPlay := calculateMyPlay(val[1], opponentPlay)
		basePoints := calculateBasePoints(myPlay)
		playPoints := calculatePlayPoints(myPlay, opponentPlay)
		pointsOutput[i] = basePoints + playPoints
	}

	return pointsOutput
}

func calculateMyPlay(myInstruction, opponentPlay string) string {
	myPlay := ""

	switch {
	// X -> lose
	case myInstruction == "X":
		switch {
		case opponentPlay == "A":
			myPlay = "Z"
		case opponentPlay == "B":
			myPlay = "X"
		case opponentPlay == "C":
			myPlay = "Y"
		}
	// Y -> draw
	case myInstruction == "Y":
		switch {
		case opponentPlay == "A":
			myPlay = "X"
		case opponentPlay == "B":
			myPlay = "Y"
		case opponentPlay == "C":
			myPlay = "Z"
		}
	// Z -> win
	case myInstruction == "Z":
		switch {
		case opponentPlay == "A":
			myPlay = "Y"
		case opponentPlay == "B":
			myPlay = "Z"
		case opponentPlay == "C":
			myPlay = "X"
		}
	}
	return myPlay
}
