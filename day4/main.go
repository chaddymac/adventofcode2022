package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// each section has IDs
// elves are assigner a range of sections
// in how many assignment pairs fully contain the other
var AssignOne []string
var AssignTwo []string
var assignments int
var assignmentOverlap int

func main() {
	fmt.Println(assignmentRanges("input.txt"))
}

func assignmentRanges(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		AssignOne = append(AssignOne, s[0])
		AssignTwo = append(AssignTwo, s[1])
	}

	for i := 0; i < len(AssignOne); i++ {
		x := strings.Split(AssignOne[i], "-")
		// y := strings.Split(AssignTwo[i], "-")
		intX, err := strconv.Atoi(x[0])
		if err != nil {
			fmt.Println(err)
		}
		intX2, err := strconv.Atoi(x[1])
		if err != nil {
			fmt.Println(err)
		}

		y := strings.Split(AssignTwo[i], "-")
		intY, err := strconv.Atoi(y[0])
		if err != nil {
			fmt.Println(err)
		}
		intY2, err := strconv.Atoi(y[1])
		if err != nil {
			fmt.Println(err)
		}
		xRange := makeRange(intX, intX2)
		yRange := makeRange(intY, intY2)

		if subset(xRange, yRange) || subset(yRange, xRange) {
			assignments += 1
		}
		if len(Intersection(xRange, yRange)) > 0 {
			assignmentOverlap += 1
		}

	}
	fmt.Println(assignmentOverlap)
	return assignments
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
