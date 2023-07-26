package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

var Stack1 Stack
var Stack2 Stack
var Stack3 Stack
var Stack4 Stack
var Stack5 Stack
var Stack6 Stack
var Stack7 Stack
var Stack8 Stack
var Stack9 Stack
var moves []string
var numChanges string
var takeFrom string
var giveTo string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) PushTop(str string) {
	*s = append([]string{str},*s...) // Simply append the new value to the end of the stack
}

func (s *Stack) Push(str string) {
	*s = append(*s,str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func setupStacks() {
	// var stack Stack // create a stack variable of type Stack

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		y := scanner.Text()
		if strings.Contains(y, "[") {
			one := string(y[0:3])
			if !strings.Contains(one, " ") {
				Stack1.PushTop(one)
			}
			two := string(y[4:7])
			if !strings.Contains(two, " ") {
				Stack2.PushTop(two)
			}
			three := string(y[8:11])
			if !strings.Contains(three, " ") {
				Stack3.PushTop(three)
			}
			four := string(y[12:15])
			if !strings.Contains(four, " ") {
				Stack4.PushTop(four)
			}
			five := string(y[16:19])
			if !strings.Contains(five, " ") {
				Stack5.PushTop(five)
			}
			six := string(y[20:23])
			if !strings.Contains(six, " ") {
				Stack6.PushTop(six)
			}
			seven := string(y[24:27])
			if !strings.Contains(seven, " ") {
				Stack7.PushTop(seven)
			}
			eight := string(y[28:31])
			if !strings.Contains(eight, " ") {
				Stack8.PushTop(eight)
			}
			nine := string(y[32:35])
			if !strings.Contains(nine, " ") {
				Stack9.PushTop(nine)
			}
		} else if strings.Contains(y, "move") {
			moves = append(moves, y)
		}
		fmt.Printf("Stack1: %v\n", Stack1)
		fmt.Printf("Stack2: %v\n", Stack2)
		fmt.Printf("Stack3: %v\n", Stack3)
		fmt.Printf("Stack4: %v\n", Stack4)
		fmt.Printf("Stack5: %v\n", Stack5)
		fmt.Printf("Stack6: %v\n", Stack6)
		fmt.Printf("Stack7: %v\n", Stack7)
		fmt.Printf("Stack8: %v\n", Stack8)
		fmt.Printf("Stack9: %v\n", Stack9)
	}
}

func makeMoves(Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9 Stack, moves []string)(Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack, Stack) {
	var chosenStack Stack
	var addStack Stack
	for _, move := range moves {
		if len(move) == 19 {
			takeFrom = string(move[13])
		} else {
			takeFrom = string(move[12])
		}
		fmt.Println(takeFrom)
		switch takeFrom {
		case "1":
			chosenStack = Stack1
		case "2":
			chosenStack = Stack2
		case "3":
			chosenStack = Stack3
		case "4":
			chosenStack = Stack4
		case "5":
			chosenStack = Stack5
		case "6":
			chosenStack = Stack6
		case "7":
			chosenStack = Stack7
		case "8":
			chosenStack = Stack8
		case "9":
			chosenStack = Stack9
		}
		if len(move) == 19 {
			giveTo = string(move[18])
		} else {
			giveTo= string(move[17])
		}

		switch giveTo {
		case "1":
			addStack = Stack1
		case "2":
			addStack = Stack2
		case "3":
			addStack = Stack3
		case "4":
			addStack = Stack4
		case "5":
			addStack = Stack5
		case "6":
			addStack = Stack6
		case "7":
			addStack = Stack7
		case "8":
			addStack = Stack8
		case "9":
			addStack = Stack9
		}
		if len(move) == 19 {
			numChanges = string(move[5:7])
		} else {
			numChanges = string(move[5])
		}
		finalNum,err := strconv.Atoi(numChanges)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < finalNum; i++ {
			x, _ := chosenStack.Pop()
			addStack.Push(x)
		}

		switch takeFrom {
		case "1":
			Stack1 = chosenStack
		case "2":
			Stack2 = chosenStack
		case "3":
			Stack3 = chosenStack
		case "4":
			Stack4 = chosenStack
		case "5":
			Stack5 = chosenStack
		case "6":
			Stack6 = chosenStack
		case "7":
			Stack7 = chosenStack
		case "8":
			Stack8 = chosenStack
		case "9":
			Stack9 = chosenStack
		}

		switch giveTo {
		case "1":
			Stack1 = addStack
		case "2":
			Stack2 = addStack
		case "3":
			Stack3 = addStack
		case "4":
			Stack4 = addStack
		case "5":
			Stack5 = addStack
		case "6":
			Stack6 = addStack
		case "7":
			Stack7 = addStack
		case "8":
			Stack8 = addStack
		case "9":
			Stack9 = addStack
		}
	}
	return Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9
}

func main() {
	setupStacks()
	fin1,fin2,fin3,fin4,fin5,fin6,fin7,fin8,fin9:= makeMoves(Stack1, Stack2, Stack3, Stack4, Stack5, Stack6, Stack7, Stack8, Stack9, moves)
	fmt.Printf("Stack 1: %v\n", fin1)
	fmt.Printf("Stack 2: %v\n", fin2)
	fmt.Printf("Stack 3: %v\n", fin3)
	fmt.Printf("Stack 4: %v\n", fin4)
	fmt.Printf("Stack 5: %v\n", fin5)
	fmt.Printf("Stack 6: %v\n", fin6)
	fmt.Printf("Stack 7: %v\n", fin7)
	fmt.Printf("Stack 8: %v\n", fin8)
	fmt.Printf("Stack 9: %v\n", fin9)
}
