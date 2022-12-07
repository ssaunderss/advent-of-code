package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputMap := DbgInput()
	inputMap := ProblemInput()
	cleanedInput := ReadInputFile("input-1.txt")
	movedCargo := CargoReducerV2(inputMap, cleanedInput)
	fmt.Println(movedCargo)
}

func ReadInputFile(fileName string) [][]int {
	f, _ := os.ReadFile(fileName)

	input := string(f)

	rows := strings.Split(input, "\n")

	sliceLength := len(rows)

	output := make([][]int, sliceLength)

	for i, val := range rows {
		splitString := strings.Split(val, " ")
		num, _ := strconv.Atoi(splitString[1])
		from, _ := strconv.Atoi(splitString[3])
		to, _ := strconv.Atoi(splitString[5])
		move := []int{num, from, to}

		output[i] = move
	}

	return output
}

/*
takes in a map representing the initial state of the cargo and a matrix representing moves
to make. Reduces the moves on the passed in state and returns the cargo (string) at the top
of each stack.
[

	[1, 2, 3],
	[3, 2, 1]

]

for each set of moves take the first n elements from input matrix int64[][] -> i[r-1][:n]:

	reverse the elements
	append to i[r2-1]
*/
func CargoReducer(baseCargo map[int][]string, moves [][]int) map[int][]string {
	for _, val := range moves {
		num, from, to := val[0], val[1], val[2]
		// fmt.Printf("Iteration %v from: %q\n", i, int(from))
		// fmt.Printf("Iteration %v num: %q\n", i, int(num))

		// pop cargo and update stack popped from
		popped := []string{}
		remaining := []string{}
		stack := baseCargo[from]
		// fmt.Printf("Iteration %v stack: %q\n", i, stack)
		popped = append(popped, stack[:num]...)
		remaining = append(remaining, stack[num:]...)
		// fmt.Printf("Iteration %v popped: %q\n", i, popped)
		// fmt.Printf("Iteration %v remaining: %q\n", i, remaining)
		baseCargo[from] = remaining
		reversePopped := ReverseSlice(popped)

		// update the cargo stack appened to
		base := baseCargo[to]
		appended := append(reversePopped, base...)
		baseCargo[to] = appended

		// fmt.Printf("Iteration %v: %q\n", i, baseCargo)
	}

	return baseCargo
}

func CargoReducerV2(baseCargo map[int][]string, moves [][]int) map[int][]string {
	for _, val := range moves {
		num, from, to := val[0], val[1], val[2]

		// pop cargo and update stack popped from
		popped := []string{}
		remaining := []string{}
		stack := baseCargo[from]
		popped = append(popped, stack[:num]...)
		remaining = append(remaining, stack[num:]...)

		baseCargo[from] = remaining

		// update the cargo stack appened to
		base := baseCargo[to]
		appended := append(popped, base...)
		baseCargo[to] = appended

		// fmt.Printf("Iteration %v: %q\n", i, baseCargo)
	}

	return baseCargo
}

func ReverseSlice(slice []string) []string {
	i := 0
	j := len(slice) - 1
	for i < j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return slice
}

func DbgInput() map[int][]string {
	inputMap := make(map[int][]string)
	inputMap[1] = []string{"N", "Z"}
	inputMap[2] = []string{"D", "C", "M"}
	inputMap[3] = []string{"P"}
	return inputMap
}

func ProblemInput() map[int][]string {
	inputMap := make(map[int][]string)
	/*
		[D]                     [N] [F]
		[H] [F]             [L] [J] [H]
		[R] [H]             [F] [V] [G] [H]
		[Z] [Q]         [Z] [W] [L] [J] [B]
		[S] [W] [H]     [B] [H] [D] [C] [M]
		[P] [R] [S] [G] [J] [J] [W] [Z] [V]
		[W] [B] [V] [F] [G] [T] [T] [T] [P]
		[Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
		 1   2   3   4   5   6   7   8   9
	*/
	inputMap[1] = []string{"D", "H", "R", "Z", "S", "P", "W", "Q"}
	inputMap[2] = []string{"F", "H", "Q", "W", "R", "B", "V"}
	inputMap[3] = []string{"H", "S", "V", "C"}
	inputMap[4] = []string{"G", "F", "H"}
	inputMap[5] = []string{"Z", "B", "J", "G", "P"}
	inputMap[6] = []string{"L", "F", "W", "H", "J", "T", "Q"}
	inputMap[7] = []string{"N", "J", "V", "L", "D", "W", "T", "Z"}
	inputMap[8] = []string{"F", "H", "G", "J", "C", "Z", "T", "D"}
	inputMap[9] = []string{"H", "B", "M", "V", "P", "W"}
	return inputMap
}
