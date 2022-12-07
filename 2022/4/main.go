package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cleanedInput := ReadInputFile("input-1.txt")
	output := calcOverlapsV2(cleanedInput)
	fmt.Println(output)
}

func ReadInputFile(fileName string) [][][]int64 {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	pairedStrings := strings.Split(input, "\n")

	sliceLength := len(pairedStrings)

	output := make([][][]int64, sliceLength)

	for i, val := range pairedStrings {
		splitString := strings.Split(val, ",")
		pairs := make([][]int64, 2)
		for j, subval := range splitString {
			splitFinal := strings.Split(subval, "-")
			start, _ := strconv.ParseInt(splitFinal[0], 0, 64)
			end, _ := strconv.ParseInt(splitFinal[1], 0, 64)
			output := []int64{start, end}
			pairs[j] = output
		}

		output[i] = pairs
	}

	return output
}

func calcOverlaps(pairs [][][]int64) int64 {
	var overlaps int64

	for _, val := range pairs {
		assignmentOne := val[0]
		assignmentTwo := val[1]

		switch {
		// equivalent to assignments being equal
		case assignmentOne[0] == assignmentTwo[0] && assignmentOne[1] == assignmentTwo[1]:
			overlaps++
		// since above didn't pass, if they share a start or end, one of them must be a sub slice
		case assignmentOne[0] == assignmentTwo[0] || assignmentOne[1] == assignmentTwo[1]:
			overlaps++
		// assignmentOne is a subslice of assignmentTwo
		case assignmentOne[0] > assignmentTwo[0] && assignmentOne[1] < assignmentTwo[1]:
			overlaps++
		// assignmentTwo is a subslice of assignmentOne
		case assignmentOne[0] < assignmentTwo[0] && assignmentOne[1] > assignmentTwo[1]:
			overlaps++
		}
	}
	return overlaps
}

func calcOverlapsV2(pairs [][][]int64) int64 {
	var overlaps int64

	for _, val := range pairs {
		assignmentOne := val[0]
		assignmentTwo := val[1]

		switch {
		// equivalent to assignments being equal
		case assignmentOne[0] == assignmentTwo[0] && assignmentOne[1] == assignmentTwo[1]:
			overlaps++
		// since above didn't pass, if they share a start or end, one of them must be a sub slice
		case assignmentOne[0] == assignmentTwo[0] || assignmentOne[1] == assignmentTwo[1]:
			overlaps++
		// assignmentOne is a subslice of assignmentTwo
		case assignmentOne[0] > assignmentTwo[0] && assignmentOne[1] < assignmentTwo[1]:
			overlaps++
		// assignmentTwo is a subslice of assignmentOne
		case assignmentOne[0] < assignmentTwo[0] && assignmentOne[1] > assignmentTwo[1]:
			overlaps++
		// partial subslice on assignmentOne tail
		case assignmentOne[1] >= assignmentTwo[0] && assignmentOne[1] <= assignmentTwo[1]:
			overlaps++
		// partial subslice on assignmentTwo tail
		case assignmentTwo[1] >= assignmentOne[0] && assignmentTwo[1] <= assignmentOne[1]:
			overlaps++
		// partial subslice on assignmentOne head
		case assignmentOne[0] >= assignmentTwo[1] && assignmentOne[0] <= assignmentTwo[0]:
			overlaps++
		// partial subslice on assignmentTwo head
		case assignmentTwo[0] >= assignmentOne[1] && assignmentTwo[0] <= assignmentOne[0]:
			overlaps++
		}
	}
	return overlaps
}
