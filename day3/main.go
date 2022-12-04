package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var compartmentOne []string
var compartmentTwo []string
var score int

// one item per rucksack is wrong
// input is the items in each rucksack
// each rucksack has the same number of items in both compartments
// first half of the input is the items in the first compartment and second half is the items in the second compartment
// 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s);

func main() {
	fmt.Println(rucksack("input.txt"))
}

func rucksack(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		y := scanner.Text()
		first := y[:len(y)/2]
		second := y[len(y)/2:]

		compartmentOne = append(compartmentOne, first)
		compartmentTwo = append(compartmentTwo, second)
	}

	for i := 0; i < len(compartmentOne); i++ {
		// utf8.RuneCountInString(str)
		if strings.Contains(compartmentTwo[i], "p") && strings.Contains(compartmentOne[i], "p") {
			score += 16
		} else if strings.Contains(compartmentTwo[i], "L") && strings.Contains(compartmentOne[i], "L") {
			score += 38
		} else if strings.Contains(compartmentTwo[i], "P") && strings.Contains(compartmentOne[i], "P") {
			score += 42
		} else if strings.Contains(compartmentTwo[i], "v") && strings.Contains(compartmentOne[i], "v") {
			score += 22
		} else if strings.Contains(compartmentTwo[i], "t") && strings.Contains(compartmentOne[i], "t") {
			score += 20
		} else if strings.Contains(compartmentTwo[i], "s") && strings.Contains(compartmentOne[i], "s") {
			score += 19
		}
	}
	return score

}
