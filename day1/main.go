package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var listCalories []string

var largestCalories int
var prevSum int
var currSum int
var allSums []int

func main() {
	fmt.Println(calorieSum())

}

func calorieSum() (int, int) {
	file, err := os.Open("reindeer_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		i := scanner.Text()
		listCalories = append(listCalories, i)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	for _, v := range listCalories {
		if v != "" {
			calorie, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			currSum += calorie
		} else {
			allSums = append(allSums, currSum)
			if prevSum > currSum && prevSum > largestCalories {
				largestCalories = prevSum
				prevSum = currSum
				currSum = 0
			} else if currSum > prevSum && currSum > largestCalories {
				largestCalories = currSum
				prevSum = currSum
				currSum = 0
			} else {
				prevSum = currSum
				currSum = 0
			}
		}

	}
	sort.Ints(allSums)

	final := allSums[len(allSums)-1] + allSums[len(allSums)-2] + allSums[len(allSums)-3]

	return largestCalories, final
}
