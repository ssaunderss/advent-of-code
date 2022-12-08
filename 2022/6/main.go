package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := ReadInputFile("input-1.txt")
	output := StartOfPacket(input, 14)
	fmt.Println(output)
}

func ReadInputFile(fileName string) string {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	rows := strings.Split(input, "\n")

	output := rows[0]

	return output
}

func StartOfPacket(dataStream string, sequenceLen int) int {
	// startingSequence := startSequence(dataStream, sequenceLen)
	output := 0
	// fmt.Println(startingSequence)
	for i, _ := range dataStream {
		if i >= sequenceLen-1 {
			sequence := dataStream[i : i+sequenceLen]
			if CheckUnique(sequence) == true {
				fmt.Println(sequence)
				return i + sequenceLen
			}
		}
	}
	return output
}

func CheckUnique(sequence string) bool {
	chars := make(map[rune]bool)
	for _, i := range sequence {
		_, ok := chars[i]
		if ok {
			return false
		}
		chars[i] = true
	}
	return true
}

// func startSequence(dataStream string, sequenceLen int) string {

// }
