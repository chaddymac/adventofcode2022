package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

var compartmentOne []string
var compartmentTwo []string
var score int
var groupOne []string
var groupTwo []string
var groupThree []string
var letter string

// one item per rucksack is wrong
// input is the items in each rucksack
// each rucksack has the same number of items in both compartments
// first half of the input is the items in the first compartment and second half is the items in the second compartment
// 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s);

func main() {
	// fmt.Println(rucksack("input.txt"))
	fmt.Println(rucksackGroup("input.txt"))
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
		val := findingMatch(compartmentOne[i], compartmentTwo[i])
		a := string(val[0])
		b, _ := utf8.DecodeRuneInString(a)
		if unicode.IsUpper(b) {
			score += int(b) - 38
		} else {
			score += int(b) - 96
		}
	}
	return score

}

func findingMatch(comp1 string, comp2 string) string {
	for i := 0; i < len(comp1); i++ {
		for j := 0; j < len(comp2); j++ {
			if comp1[i] == comp2[j] {
				return comp1[i:]
			}
		}
	}
	return ""
}

func rucksackGroup(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		y := scanner.Text()

		groupOne = append(groupOne, y)
	}

	for i := 0; i < len(groupOne); i += 3 {
		firstItem := groupOne[i]
		secondItem := groupOne[i+1]
		thirdItem := groupOne[i+2]
		letter = findMatchTwo(firstItem, secondItem, thirdItem)
		b, _ := utf8.DecodeRuneInString(letter)
		if unicode.IsUpper(b) {
			score += int(b) - 38
		} else {
			score += int(b) - 96
		}

	}
	return score
}

func findMatchTwo(comp1 string, comp2 string, comp3 string) string {
	for i := 0; i < len(comp1); i++ {
		for j := 0; j < len(comp2); j++ {
			for k := 0; k < len(comp3); k++ {
				if comp1[i] == comp2[j] && comp1[i] == comp3[k] {
					letter = string(comp1[i])
					return letter
				}
			}
		}
	}
	return ""
}
