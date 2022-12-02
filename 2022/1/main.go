package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cleanedSlice := readInputFile("input-2.txt")
	answer := trackLargestV2(cleanedSlice)
	fmt.Printf("Answer: %d\n", answer)
}

func readInputFile(fileName string) []int64 {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	someStrings := strings.Split(input, "\n")

	var someNums = []int64{}

	/*
		Parse strings to int and stuff into a slice,
		if can't (because it's blank), just put -1
	*/
	for _, val := range someStrings {
		num, err := strconv.ParseInt(val, 0, 64)
		if err != nil {
			someNums = append(someNums, -1)
		} else {
			fmt.Printf("%d\n", num)
			someNums = append(someNums, num)
		}
	}

	return someNums
}

func trackLargest(calories []int64) int64 {
	var largestNum int64 = 0
	var runningCalorieCount int64 = 0
	lastIndex := len(calories) - 1

	for i, value := range calories {
		switch {
		case i == 1:
			runningCalorieCount = value
		case value == -1 || i == lastIndex:
			if runningCalorieCount > largestNum {
				largestNum = runningCalorieCount
			}
			runningCalorieCount = 0
		default:
			runningCalorieCount += value
		}
	}

	return largestNum
}

func trackLargestV2(calories []int64) int64 {
	var largestNums = []int64{0, 0, 0}
	var runningCalorieCount int64 = 0
	var sum int64 = 0
	lastIndex := len(calories) - 1

	for i, value := range calories {
		switch {
		case i == 1:
			runningCalorieCount = value
		case value == -1 || i == lastIndex:
			updateLargestNums(runningCalorieCount, largestNums)
			runningCalorieCount = 0
		default:
			runningCalorieCount += value
		}
	}

	for _, val := range largestNums {
		sum += val
	}

	return sum
}

func updateLargestNums(calorieCount int64, largestNums []int64) []int64 {
	set := false
	for i, value := range largestNums {
		if calorieCount > value && set != true {
			largestNums[i] = calorieCount
			set = true
		}
	}
	return largestNums
}
