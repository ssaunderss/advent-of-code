package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cleanedSlice := ReadInputFile("input-1.txt", 2)
	// rawPoints := CalculatePoints(cleanedSlice)
	rawPoints := CalculatePointsV2(cleanedSlice)
	output := AggregatePoints(rawPoints)
	fmt.Println(output)
}

func ReadInputFile(fileName string, version int) [][]string {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	pairedStrings := strings.Split(input, "\n")

	sliceLength := len(pairedStrings)

	output := make([][]string, sliceLength)

	/*
		Creates matrix with each elf's compartments
	*/
	if version == 1 {
		for i, val := range pairedStrings {
			splitString := SplitMeDownTheMiddle(val)
			output[i] = splitString
		}
	} else if version == 2 {
		output = ChunkMe(pairedStrings, 3)
	}

	return output
}

func CalculatePoints(compartmentMatrix [][]string) []int64 {
	outputLen := len(compartmentMatrix)
	output := make([]int64, outputLen)

	/*
		Wouldn't be able to use this approach if we weren't exclusively
		using latin alphabet because range takes runes instead of full chars
	*/

	for i, val := range compartmentMatrix {
		left := val[0]
		right := val[1]
		matchingChar := ""

		for _, rune := range left {
			char := string(rune)
			if strings.Contains(right, char) {
				matchingChar = char
				break
			}
		}
		// lookup matching char and append
		priority := LookupCharPriority(matchingChar)
		output[i] = priority
	}

	return output
}

func CalculatePointsV2(groups [][]string) []int64 {
	outputLen := len(groups)
	output := make([]int64, outputLen)

	/*
		Wouldn't be able to use this approach if we weren't exclusively
		using latin alphabet because range takes runes instead of full chars
	*/

	for i, val := range groups {
		first := val[0]
		second := val[1]
		third := val[2]
		matchingChar := ""

		for _, rune := range first {
			char := string(rune)
			if strings.Contains(second, char) && strings.Contains(third, char) {
				matchingChar = char
				break
			}
		}
		// lookup matching char and append
		priority := LookupCharPriority(matchingChar)
		output[i] = priority
	}

	return output
}

func AggregatePoints(rawPoints []int64) int64 {
	var totalPoints int64 = 0
	for _, val := range rawPoints {
		totalPoints += val
	}
	return totalPoints
}

func SplitMeDownTheMiddle(inputString string) []string {
	stringLen := len(inputString)
	// isEven := (stringLen%2 == 0)
	// fmt.Printf("Input string is even?: %t\n", isEven)

	halfwayPoint := stringLen / 2
	output := make([]string, 2)
	output[0] = inputString[:halfwayPoint]
	output[1] = inputString[halfwayPoint:]
	return output
}

func ChunkMe(input []string, chunkSize int) [][]string {
	outputLen := len(input) / chunkSize
	output := make([][]string, outputLen)

	for i := 0; i < outputLen; i++ {
		start := i * chunkSize
		end := start + chunkSize
		chunk := input[start:end]
		output[i] = chunk
	}

	return output
}

func LookupCharPriority(char string) int64 {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var priority int64 = 0

	for i, val := range letters {
		if char == string(val) {
			priority = int64(i) + 1
			break
		}
	}
	return priority
}
