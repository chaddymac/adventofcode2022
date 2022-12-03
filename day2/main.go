package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var m map[string]int
var opponent []string
var choice []string
var score int
var newScore int

// Opponent A for rock, B for paper, C for scissors.
// You X for rock, Y for paper, Z for scissors.
// 1 point for rock, 2 points for paper, 3 points for scissors.
// 0 for loss, 3 for draw and 6 for win.
//X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

// for A need Y for draw 3 and 1, B need X for loss 1 and 0, C need Z for win 3 and 9
func main() {
	totalScore()
}

func totalScore() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		// fmt.Println(s)
		opponent = append(opponent, s[0])
		choice = append(choice, s[1])
	}

	for i := 0; i < len(choice); i++ {
		switch choice[i] {
		case "X":
			if opponent[i] == "A" {
				score += 3
			} else if opponent[i] == "B" {
				score += 0
			} else {
				score += 6
			}
			score += 1
		case "Y":
			if opponent[i] == "A" {
				score += 6
			} else if opponent[i] == "B" {
				score += 3
			} else {
				score += 0
			}
			score += 2
		case "Z":
			if opponent[i] == "A" {
				score += 0
			} else if opponent[i] == "B" {
				score += 6
			} else {
				score += 3
			}
			score += 3
		}

		switch opponent[i] {
		case "A":
			if choice[i] == "Y" {
				newScore += 3
				newScore += 1
			} else if choice[i] == "X" {
				newScore += 0
				newScore += 3
			} else {
				newScore += 6
				newScore += 2
			}
		case "B":
			if choice[i] == "X" {
				newScore += 0
				newScore += 1
			} else if choice[i] == "Y" {
				newScore += 3
				newScore += 2
			} else {
				newScore += 6
				newScore += 3
			}
		case "C":
			if choice[i] == "Z" {
				newScore += 6
				newScore += 1
			} else if choice[i] == "Y" {
				newScore += 3
				newScore += 3
			} else {
				newScore += 0
				newScore += 2
			}
		}
	}
	fmt.Println(score)
	fmt.Println(newScore)
	return score, scanner.Err()
}
